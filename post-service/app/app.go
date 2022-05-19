package app

import (
	"github.com/bektosh03/test-crud/post-service/app/command"
	"github.com/bektosh03/test-crud/post-service/app/query"
	"github.com/bektosh03/test-crud/post-service/domain/post"
)

type App struct {
	Commands Commands
	Queries  Queries
}

func New(postRepo post.Repository) App {
	return App{
		Commands: Commands{
			UpdatePost: command.NewUpdatePostHandler(postRepo),
			DeletePost: command.NewDeletePostHandler(postRepo),
		},
		Queries: Queries{
			GetPost:   query.NewGetPostHandler(postRepo),
			ListPosts: query.NewListPostsHandler(postRepo),
		},
	}
}

type Commands struct {
	UpdatePost command.UpdatePostHandler
	DeletePost command.DeletePostHandler
}

type Queries struct {
	GetPost   query.GetPostHandler
	ListPosts query.ListPostsHandler
}
