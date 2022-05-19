package ports

import (
	"context"
	"fmt"

	"github.com/bektosh03/test-crud/data-service/app"
	"github.com/bektosh03/test-crud/common/errs"
	"github.com/bektosh03/test-crud/genprotos/datapb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	app *app.App
	datapb.UnimplementedDataServiceServer
}

func NewGrpcServer(app *app.App) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (s GrpcServer) DownloadPosts(ctx context.Context, _ *datapb.DownloadPostsRequest) (*emptypb.Empty, error) {
	success, _, err := s.app.GetDownloadStatus(ctx)
	if err != nil && err != errs.ErrNotFound {
		return &emptypb.Empty{}, status.Error(codes.Internal, fmt.Sprintf("counld not get download status: %v", err))
	}

	if success {
		return &emptypb.Empty{}, status.Error(codes.AlreadyExists, "posts are already fetched")
	}

	go func() {
		err := s.app.DownloadPosts(ctx)
		if err != nil {
			logrus.Info("setting operations status to unsuccessfull")
			if updateErr := s.app.SetDownloadStatus(ctx, false, err); updateErr != nil {
				logrus.Errorf("failed to set download status: %v", updateErr)
			}
		}
	}()

	return &emptypb.Empty{}, nil
}

func (s GrpcServer) GetDownloadStatus(ctx context.Context, _ *datapb.GetDownloadStatusRequest) (*datapb.GetDownloadStatusResponse, error) {
	success, errMsg, err := s.app.GetDownloadStatus(ctx)
	if err != nil && err != errs.ErrNotFound {
		return &datapb.GetDownloadStatusResponse{}, status.Error(codes.Internal, err.Error())
	} else if err == errs.ErrNotFound {
		return &datapb.GetDownloadStatusResponse{
			Success: false,
			ErrMsg:  "download request has not yet been received",
		}, nil
	}

	return &datapb.GetDownloadStatusResponse{
		Success: success,
		ErrMsg:  errMsg,
	}, nil
}
