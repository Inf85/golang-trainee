package main

import (
	"./db-connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const postsURI string = "https://jsonplaceholder.typicode.com/posts"
const commentsURI string = "https://jsonplaceholder.typicode.com/comments"
const userID string = "7"

/**
Post
*/
type Post struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

/**
Comments
*/
type Comments struct {
	ID     int
	PostID int
	Name   string
	Email  string
	Body   string
}

var c chan []Comments = make(chan []Comments)

func main() {
	posts := getPosts()
	savePosts(posts)
	for i := range posts{
		go getComments(posts[i].ID, c)
		go saveComments(c)
	}

	var input string
	fmt.Println("Press any key ...")
	fmt.Scanln(&input)
}

/**
Get Posts
*/
func getPosts() []Post {
	var post []Post
	response, err := http.Get(postsURI + "?userId=" + userID)

	if err != nil {
		panic(err)
		return nil
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	//Check For Errors
	if error != nil {
		panic(error)
		return nil
	}
	errJSON := json.Unmarshal(body, &post)
	if err != nil {
		panic(errJSON)
		return nil
	}

	return post
}

/**
Save Post to DataBase
*/
func savePosts(posts []Post) {
	db := db_connect.SetDBConnection()

	for i := range posts {
		var bodyText string = strings.ReplaceAll(posts[i].Body, "\n", "")
		values := map[string]interface{}{
			"id":     posts[i].ID,
			"userId": posts[i].UserID,
			"title":  posts[i].Title,
			"body":   bodyText,
		}

		db.Exec("INSERT into posts (id,user_id,title,body) VALUES (@id,@userId,@title, @body)"+
			" ON DUPLICATE KEY UPDATE user_id=@userId, title=@title, body=@body",
			values)

	}

}

/**
Get Comments
*/
func getComments(postId int, c chan<- []Comments) {
	var comments []Comments
	var strPostId string = strconv.Itoa(postId)

	response, err := http.Get(commentsURI + "?postId=" + strPostId)

	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	//Check For Errors
	if error != nil {
		panic(error)

	}
	errJson := json.Unmarshal(body, &comments)
	if err != nil {
		panic(errJson)

	}

	c <- comments
}

/**
Save Comment to Data Base
*/
func saveComments(c <-chan []Comments) {
	var comments []Comments

	db := db_connect.SetDBConnection()

	comments = <-c

	for i := range comments {

		values := map[string]interface{}{
			"id":     comments[i].ID,
			"postId": comments[i].PostID,
			"email":  comments[i].Email,
			"name":   comments[i].Name,
			"body":   comments[i].Body,
		}

		db.Exec("INSERT into comments (id,post_id,email,name,body) VALUES (@id,@postId,@email,@name,@body)"+
			"ON DUPLICATE KEY UPDATE post_id=@postId, email=@email, name=@name, body=@body",
			values)

	}
}
