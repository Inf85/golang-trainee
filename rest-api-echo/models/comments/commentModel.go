package comments

/*
Comments Struct
*/
type Comments struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
