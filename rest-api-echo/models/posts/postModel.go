package posts

import "../../database"

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