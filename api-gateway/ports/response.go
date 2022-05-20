package ports

import "github.com/bektosh03/test-crud/genprotos/postpb"

type ResponseMessage struct {
	Message string `json:"message"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ListPostsResponse struct {
	Posts []Post `json:"posts"`
}

func toListPostsResponse(posts []*postpb.Post) ListPostsResponse {
	postsList := make([]Post, 0, len(posts))
	for _, p := range posts {
		postsList = append(postsList, Post{
			ID:     int(p.Id),
			UserID: int(p.UserId),
			Title:  p.Title,
			Body:   p.Body,
		})
	}

	return ListPostsResponse{
		Posts: postsList,
	}
}
