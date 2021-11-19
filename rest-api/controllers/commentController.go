package controllers

import (
	"encoding/json"
	"net/http"
	"../models/comment"
)

/*
CommentController - struct
 */
type CommentController struct {}

/*
GetAllComments - Get All Comments
 */
func (CommentController) GetAllComments(typeResponse string)http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var commentModel = comment.Comments{}

		res, err := commentModel.GetAll()

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