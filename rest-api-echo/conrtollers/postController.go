package conrtollers

import (
	"github.com/labstack/echo"
	"net/http"
)
import "../models/posts"
import "../helpers/response"

/*
PostController - Post Controller
 */
type PostController struct {}

/*
GetAllPosts - Get All Posts From DataBase
 */
func (postController *PostController)GetAllPosts(typeResponse string) echo.HandlerFunc  {
	return func(context echo.Context) error {
		postModel := &posts.Posts{}

		result, error := postModel.GetAll()

		if error != nil{
			panic(error)
		}

		if len(result) > 0 {
			responsehelper.XMLJSONResponseHelper(result, typeResponse, context)
		}else{
			return context.NoContent(http.StatusNoContent)
		}

		return nil
	}
}

/*
GetByID - Get Post Record By ID
 */
func (postController *PostController) GetByID(typeResponse string) echo.HandlerFunc  {
	return func(context echo.Context) error {
		postModel := &posts.Posts{}
		id := context.Param("id")

		result, error := postModel.GetPostByID(id)

		if error != nil{
			panic(error)
		}

		if len(result) > 0 {
			responsehelper.XMLJSONResponseHelper(result, typeResponse, context)
		}else{
			return context.NoContent(http.StatusNoContent)
		}

		return nil
	}

}