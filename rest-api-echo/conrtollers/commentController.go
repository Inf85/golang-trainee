package conrtollers

import (
	"../helpers/response"
	"../models/comments"
	context2 "context"
	"github.com/labstack/echo"
	http "net/http"
	"strconv"
)

/*
CommentController - commentController Struct
*/
type CommentController struct {
	service comments.Storage
}

/*
NewCommentController - Constructor
*/
func NewCommentController(service comments.Storage) *CommentController {
	return &CommentController{service: service}
}

/*
GetAllComments - Get All Comments From DataBase
*/
// @Tags Comments
// @Summary ListComments godoc
// @Summary Show all Comments
// @Description get all Comments
// @Accept  json
// @Produce  json
// @Produce  xml
// @Router /json/comments [get]
// @Router /xml/comments [get]
// @Success 200 {array} posts.Posts
// @Failure 500
func (commentController *CommentController) GetAllComments(context echo.Context) error {
	typeResponse := context.Request().Header.Get("Accept")
	result, error := commentController.service.GetAll(context2.Background(), []comments.Comments{})

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

/*
GetCommentByID - Get Commetn By ID Controller Method
*/
// @Tags Comments
// @Summary Comment Record By ID godoc
// @Summary Show Comment Record By ID
// @Description get  Comment record By ID
// @Accept  application/json
// @Accept application/xml
// @Router /comments/{id} [get]
// @Param id query string false "id"
// @Success 200 {array} comments.Comments
// @Failure 500
func (commentController *CommentController) GetCommentByID(context echo.Context) error {
	typeResponse := context.Request().Header.Get("Accept")
	id := context.Param("id")
	result, err := commentController.service.GetByID(context2.Background(), []comments.Comments{}, id)

	if err != nil {
		panic(err)
	}

	if len(result) > 0 {
		responsehelper.XMLJSONResponseHelper(result, typeResponse, context)
	} else {
		return context.NoContent(http.StatusNoContent)
	}

	return nil
}

/*
Create - Create Comment Record in database
*/
// @Tags Comments
// @Success 200 {string} string OK
// @Summary  Create Comment Record godoc
// @Summary Create Comment Record
// @Description Create Comment Record
// @Accept  json
// @Produce  json
// @Router /comment [post]
// @Param postId formData string false "postId"
// @Param email formData string false "email"
// @Param name formData string false "name"
// @Param body formData string false "body"
// @Failure 500
func (commentController *CommentController) Create(context echo.Context) error {
	data := context.Request()
	postID, _ := strconv.Atoi(data.FormValue("postId"))
	email := data.FormValue("email")
	name := data.FormValue("name")
	body := data.FormValue("body")

	insertData := comments.Comments{
		PostID: postID,
		Name:   name,
		Email:  email,
		Body:   body,
	}

	commentController.service.CreateComment(context2.Background(), &insertData)

	return context.String(http.StatusOK, "OK")
}

/*
DeleteByID - Delete Comment Record Controller Method
*/
// @Tags Comments
// @Summary  Delete Comment Record By ID godoc
// @Summary Delete Comment Record By ID
// @Description Delete Comment Record By ID
// @Accept  json
// @Produce  json
// @Router /comment [delete]
// @Param id query string false "id"
// @Success 200 {string} string OK
// @Failure 500
func (commentController *CommentController) DeleteByID(context echo.Context) error {
	id := context.QueryParam("id")
	_, err := commentController.service.DeleteByID(context2.Background(), &comments.Comments{}, id)
	if err != nil {
		panic(err)
	}

	return context.String(http.StatusOK, "OK")
}
