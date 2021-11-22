package comments

import (
	"../../database"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

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

/*
GetAll - Get All Comments From DataBase
*/
func (commentsModel Comments) GetAll() ([]Comments, error) {
	var comments []Comments

	db := dbconnect.SetDBConnection()
	db.Find(&comments)

	return comments, nil
}

/*
GetByID - Get Comment Record By ID
 */
func (commentsModel Comments) GetByID(id string) ([]Comments, error)  {
	var comment []Comments

	db := dbconnect.SetDBConnection()
	db.Find(&comment, id)

	return comment, nil
}

/*
Create - Create Comment Record
 */
func (commentsModel Comments) Create(data *http.Request)  (*gorm.DB, error) {
	db := dbconnect.SetDBConnection()

	postID, _ := strconv.Atoi(data.FormValue("postId"))
	email := data.FormValue("email")
	name := data.FormValue("name")
	body := data.FormValue("body")

	commentData := Comments{PostID: postID, Email: email, Name: name, Body: body}

	result := db.Create(&commentData)

	return result, nil
}

/*
DeleteByID - delete comment record from DataBase
 */
func (commentsModel Comments) DeleteByID(id string) (*gorm.DB, error)  {
	var comment []Comments

	db := dbconnect.SetDBConnection()

	result := db.Delete(&comment, id)

	return result, nil
}