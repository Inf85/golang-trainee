package controllers

import (
	"../helpers"
	"../models/post"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

/*
PostController struct
 */
type PostController struct {

}

/*
GetAllPosts - Get All Posts
 */
func (PostController)GetAllPosts(typeResponse string) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		var postModel = post.Posts{}

		res, err := postModel.GetAll()

		if err != nil {
			panic(err)
		}

	   helpers.TypeResponse(res, typeResponse, writer)

	})
}


/*
GetPostByID - GetPost By ID
 */
func (PostController)GetPostByID(typeResponse string)  http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var postModel = post.Posts{}

		vars := mux.Vars(request)
		 id  := vars["id"]

		res, err := postModel.GetByID(id)
		if err != nil {
			panic(err)
		}

		helpers.TypeResponse(res, typeResponse, writer)

	})

}

/*
CreatePost - Create Post
 */
func (PostController)CreatePost(writer http.ResponseWriter, request *http.Request)  {
	var postModel = post.Posts{}
	result, err := postModel.Create(request)

	if err != nil{
		var errMsg map[string]string = map[string]string{"Result":"Error", "Message":"Record Not Added"}
		resultJSON, _ := json.Marshal(errMsg)
		writer.Write(resultJSON)
	}else{
		resultJSON, _ := json.Marshal(result)
		writer.Write(resultJSON)
	}


}

/*
DeletePostByID - Delete Post
 */
func (PostController)DeletePostByID(writer http.ResponseWriter, request *http.Request) {
		var postModel = post.Posts{}
		vars := mux.Vars(request)
		id  := vars["id"]

		_, err := postModel.DeleteByID(id)

		if err != nil{
			var errMsg map[string]string = map[string]string{"Result":"Error", "Message":"Record Not Added"}
			resultJSON, _ := json.Marshal(errMsg)
			writer.Write(resultJSON)
		}else{
			var successMsg map[string]string = map[string]string{"Result":"Success", "Message":"Record Deleted"}
			resultJSON, _ := json.Marshal(successMsg)
			writer.Write(resultJSON)
		}
}