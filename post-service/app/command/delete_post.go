package command

import (
	"context"

	"github.com/bektosh03/test-crud/post-service/domain/post"
)

type DeletePostHandler struct {
	postRepo post.Repository
}

func NewDeletePostHandler(postRepo post.Repository) DeletePostHandler {
	return DeletePostHandler{
		postRepo: postRepo,
	}
}

func (h DeletePostHandler) Handle(ctx context.Context, postID int) error {
	return h.postRepo.DeletePost(ctx, postID)
}
