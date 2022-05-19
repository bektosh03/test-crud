package command

import (
	"context"

	"github.com/bektosh03/test-crud/post-service/domain/post"
)

type UpdatePostHandler struct {
	postRepo post.Repository
}

func NewUpdatePostHandler(postRepo post.Repository) UpdatePostHandler {
	return UpdatePostHandler{
		postRepo: postRepo,
	}
}

func (h UpdatePostHandler) Handle(ctx context.Context, p post.Post) error {
	return h.postRepo.UpdatePost(ctx, p)
}
