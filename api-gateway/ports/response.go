package ports

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
