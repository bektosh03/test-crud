package post

import "context"

type Repository interface {
	CreatePosts(ctx context.Context, postsBatch <-chan []Post) error
}
