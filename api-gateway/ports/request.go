package ports

type GetPostRequest struct {
	PostID int `json:"post_id"`
}

type UpdatePostRequest struct {
	PostID int    `json:"post_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type DeletePostRequest struct {
	PostID int `json:"post_id"`
}
