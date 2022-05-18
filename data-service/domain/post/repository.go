package post

import "context"

type Repository interface {
	CreatePostsAsync(ctx context.Context, postsBatch <-chan []Post) <-chan error
}
