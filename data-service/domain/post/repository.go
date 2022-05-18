package post

import "context"

type Repository interface {
	CreatePostsAsync(ctx context.Context, postsBatch <-chan []Post) <-chan error
	SetDownloadStatus(ctx context.Context, success bool, downloadErr error) error
	GetDownloadStatus(ctx context.Context) (success bool, errMsg string, err error)
}
