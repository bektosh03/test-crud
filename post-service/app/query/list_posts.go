package query

import (
	"context"

	"github.com/bektosh03/test-crud/post-service/domain/post"
)

type ListPostsHandler struct {
	postRepo post.Repository
}

func NewListPostsHandler(postRepo post.Repository) ListPostsHandler {
	return ListPostsHandler{
		postRepo: postRepo,
	}
}

func (h ListPostsHandler) Handle(ctx context.Context, page, limit int) ([]post.Post, error) {
	if page < 1 {
		page = 1
	}
	if limit < 10 {
		limit = 10
	}
	
	return h.postRepo.ListPosts(ctx, page, limit)
}
