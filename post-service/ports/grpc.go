package ports

import (
	"context"
	"errors"

	"github.com/bektosh03/test-crud/common/errs"
	"github.com/bektosh03/test-crud/genprotos/postpb"
	"github.com/bektosh03/test-crud/post-service/app"
	"github.com/bektosh03/test-crud/post-service/domain/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	app app.App
	postpb.UnimplementedPostServiceServer
}

func NewGrpcServer(app app.App) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (s GrpcServer) GetPost(ctx context.Context, request *postpb.GetPostRequest) (*postpb.Post, error) {
	p, err := s.app.Queries.GetPost.Handle(ctx, int(request.PostId))
	if errors.Is(err, errs.ErrNotFound) {
		return &postpb.Post{}, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return &postpb.Post{}, status.Error(codes.Internal, err.Error())
	}

	return toProtoPost(p), nil
}

func (s GrpcServer) ListPosts(ctx context.Context, request *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	posts, err := s.app.Queries.ListPosts.Handle(ctx, int(request.Page), int(request.Limit))
	if err != nil {
		return &postpb.ListPostsResponse{}, status.Error(codes.Internal, err.Error())
	}

	protoPosts := make([]*postpb.Post, 0, len(posts))
	for _, p := range posts {
		protoPosts = append(protoPosts, toProtoPost(p))
	}

	return &postpb.ListPostsResponse{
		Posts: protoPosts,
	}, nil
}

func (s GrpcServer) UpdatePost(ctx context.Context, request *postpb.Post) (*emptypb.Empty, error) {
	p, err := post.New(int(request.Id), int(request.UserId), request.Title, request.Body)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.app.Commands.UpdatePost.Handle(ctx, p)
	if errors.Is(err, errs.ErrNotFound) {
		return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s GrpcServer) DeletePost(ctx context.Context, request *postpb.DeletePostRequest) (*emptypb.Empty, error) {
	err := s.app.Commands.DeletePost.Handle(ctx, int(request.PostId))
	if errors.Is(err, errs.ErrNotFound) {
		return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func toProtoPost(p post.Post) *postpb.Post {
	return &postpb.Post{
		Id:     int64(p.ID()),
		UserId: int64(p.UserID()),
		Title:  p.Title(),
		Body:   p.Body(),
	}
}
