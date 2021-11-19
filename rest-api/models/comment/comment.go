package comment

import "../../database"

/*
Comments Struct
*/
type Comments struct {
	ID int `json:"id"`
	PostID int `json:"post_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string `json:"body"`
}

/*
GetAll - Get All Comments From DataBase
 */
func (commentsModel Comments) GetAll() ([]Comments, error) {
	var comments []Comments

	db:=dbconnect.SetDBConnection()
	db.Find(&comments)

	return comments,nil
}
