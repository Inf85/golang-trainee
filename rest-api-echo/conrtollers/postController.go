package conrtollers

import (
	"github.com/labstack/echo"
	"net/http"
)
import "../models/posts"
import "../helpers/response"
import _ "../docs"

/*
PostController - Post Controller
*/
type PostController struct{}

/*
NewPostController - Constructor
*/
func NewPostController() *PostController {
	return &PostController{}
}

/*
GetAllPosts - Get All Posts From DataBase
*/
// @Tags Posts
// @Summary ListPosts godoc
// @Summary Show all Post in json
// @Description get all Posts records
// @Accept  json
// @Produce  json
// @Produce  xml
// @Router /json/posts [get]
// @Router /xml/posts [get]
// @Success 200 {array} posts.Posts
// @Failure 500
func (postController *PostController) GetAllPosts(typeResponse string) echo.HandlerFunc {
	return func(context echo.Context) error {
		postModel := &posts.Posts{}

		result, error := postModel.GetAll()

		if error != nil {
			panic(error)
		}

		if len(result) > 0 {
			return responsehelper.XMLJSONResponseHelper(result, typeResponse, context)
		} else {
			return context.NoContent(http.StatusNoContent)
		}

		return nil
	}
}

/*
GetByID - Get Post Record By ID
*/
// @Tags Posts
// @Summary Post Record By ID godoc
// @Summary Show Post Record By ID
// @Description get  Post record By ID
// @Accept  json
// @Produce  json
// @Produce  xml
// @Router /json/posts/{id} [get]
// @Router /xml/posts/{id} [get]
// @Param id query string false "id"
// @Success 200 {array} posts.Posts
// @Failure 500
func (postController *PostController) GetByID(typeResponse string) echo.HandlerFunc {
	return func(context echo.Context) error {
		postModel := &posts.Posts{}
		id := context.Param("id")

		result, error := postModel.GetPostByID(id)

		if error != nil {
			panic(error)
		}

		if len(result) > 0 {
			responsehelper.XMLJSONResponseHelper(result, typeResponse, context)
		} else {
			return context.NoContent(http.StatusNoContent)
		}

		return nil
	}

}

/*
Create - Create Record Controller Method
*/
// @Tags Posts
// @Success 200 {string} string OK
// @Summary  Create Post Record godoc
// @Summary Create Post Record
// @Description Create Post Record
// @Accept  json
// @Produce  json
// @Router /post [post]
// @Param userId formData string false "userId"
// @Param title formData string false "title"
// @Param body formData string false "body"
// @Failure 500
func (postController *PostController) Create(context echo.Context) error {
	postModel := &posts.Posts{}
	data := context.Request()

	_, error := postModel.Create(data)

	if error != nil {
		return error
	}
	return context.String(http.StatusOK, "OK")

}

/*
Delete - Delete Record By ID Controller Method
*/
// @Tags Posts
// @Summary  Delete Post Record By ID godoc
// @Summary Delete Post Record By ID
// @Description Delete Post Record By ID
// @Accept  json
// @Produce  json
// @Router /post [delete]
// @Param id query string false "id"
// @Success 200 {string} string OK
// @Failure 500
func (postController *PostController) Delete(context echo.Context) error {
	postModel := &posts.Posts{}
	id := context.QueryParam("id")

	_, error := postModel.DeleteByID(id)

	if error != nil {
		return error
	}
	return context.String(http.StatusOK, "OK")
}
