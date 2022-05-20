package ports

type UpdatePostRequest struct {
	PostID int    `json:"post_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
