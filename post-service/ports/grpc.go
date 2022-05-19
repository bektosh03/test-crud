package ports

import (
	"context"

	"github.com/bektosh03/test-crud/genprotos/postpb"
	"github.com/bektosh03/test-crud/post-service/app"
	"github.com/bektosh03/test-crud/post-service/domain/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.App
}

func NewGrpcServer(app app.App) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (s GrpcServer) GetPost(ctx context.Context, request *postpb.GetPostRequest) (*postpb.Post, error) {
	p, err := s.app.Queries.GetPost.Handle(ctx, int(request.PostId))
	if err != nil {
		return &postpb.Post{}, status.Error(codes.Internal, err.Error())
	}

	return toProtoPost(p), nil
}

func toProtoPost(p post.Post) *postpb.Post {
	return &postpb.Post{
		Id:     int64(p.ID()),
		UserId: int64(p.UserID()),
		Title:  p.Title(),
		Body:   p.Body(),
	}
}
