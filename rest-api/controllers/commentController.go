package controllers

import (
	"../helpers"
	"../models/comment"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

/*
CommentController - struct
*/
type CommentController struct{}

/*
GetAllComments - Get All Comments
*/
func (CommentController) GetAllComments(typeResponse string) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var commentModel = comment.Comments{}

		res, err := commentModel.GetAll()

		if err != nil {
			panic(err)
		}

		helpers.TypeResponse(res, typeResponse, writer)
	})
}

/*
GetCommentByID - Get Comment By ID
 */
func (CommentController) GetCommentByID(typeResponse string) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var commentModel = comment.Comments{}
		vars := mux.Vars(request)
		id := vars["id"]

		res, err := commentModel.GetByID(id)

		if err != nil {
			panic(err)
		}
		helpers.TypeResponse(res, typeResponse, writer)
	})

}

/*
CreateComment - Create comment
 */
func (CommentController) CreateComment(writer http.ResponseWriter, request *http.Request)  {
	var commentModel = comment.Comments{}

	result, err := commentModel.Create(request)

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
DeleteCommentByID - Delete Comment From DataBAse
 */
func (CommentController) DeleteCommentByID(writer http.ResponseWriter, request *http.Request) {
		var commentModel = comment.Comments{}
		vars := mux.Vars(request)
		id := vars["id"]

		_, err := commentModel.DeleteByID(id)

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
