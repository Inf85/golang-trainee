package controllers

import (
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

		if typeResponse == "json" {
			resultJSON, _ := json.Marshal(res)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(resultJSON)
		}
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

		if typeResponse == "json" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)

			if  len(res) == 0{
				var errMsg map[string]string = map[string]string{"Result":"Error", "Message":"Record Not Found"}

				resultJSON, _ := json.Marshal(errMsg)
				writer.Write(resultJSON)
			}else {
				resultJSON, _ := json.Marshal(res)
				writer.Write(resultJSON)
			}
		}

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