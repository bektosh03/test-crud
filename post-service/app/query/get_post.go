package query

import (
	"context"

	"github.com/bektosh03/test-crud/post-service/domain/post"
)

type GetPostHandler struct {
	postRepo post.Repository
}

func NewGetPostHandler(postRepo post.Repository) GetPostHandler {
	return GetPostHandler{
		postRepo: postRepo,
	}
}

func (h GetPostHandler) Handle(ctx context.Context, postID int) (post.Post, error) {
	return h.postRepo.GetPost(ctx, postID)
}
