package post

import (
	"context"
)

type Repository interface {
	GetPost(ctx context.Context, postID int) (Post, error)
	ListPosts(ctx context.Context, page, limit int) ([]Post, error)
	UpdatePost(ctx context.Context, p Post) error
	DeletePost(ctx context.Context, postID int) error
}
