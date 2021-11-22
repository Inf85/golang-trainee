package posts

import (
	"../../database"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

/*
Posts Struct
*/
type Posts struct {
	UserID int `json:"user_id"`
	ID int `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Body string `json:"body"`
	Comments []Comments `json:"comment" gorm:"foreignKey:post_id"`
}

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
GetAll - Get Posts Records
*/
func (postModel *Posts) GetAll() ([]Posts, error) {
	var posts []Posts

	db := dbconnect.SetDBConnection()
	db.Preload("Comments").Find(&posts)

	return posts, nil
}

/*
GetPostByID - GEt Post By Id From DataBase
 */
func (postModel *Posts) GetPostByID(id string) ([]Posts, error) {
	var post []Posts

	db := dbconnect.SetDBConnection()
	db.Preload("Comments").Find(&post, id)

	return post, nil
}

/*
Create - Create Record in database
 */
func (postModel *Posts) Create(data *http.Request) (*gorm.DB, error) {
	db := dbconnect.SetDBConnection()

	userID, _ := strconv.Atoi(data.FormValue("userId"))
	title := data.FormValue("title")
	body := data.FormValue("body")

	postData := Posts{UserID: userID, Title: title, Body: body}

	result := db.Create(&postData)

	return result, nil
}

/*
DeleteByID - Delete Record from database by ID
 */
func (postModel * Posts) DeleteByID(id string) (*gorm.DB, error)  {
	var post []Posts
	db := dbconnect.SetDBConnection()

	result := db.Delete(&post, id)

	return result, nil
}