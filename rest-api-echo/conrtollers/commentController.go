package conrtollers

import (
	"../models/comments"
	"../helpers/response"
	"github.com/labstack/echo"
	http "net/http"
)

/*
CommentController - commentController Struct
 */
type CommentController struct {}

/*
NewCommentController - Constructor
 */
func NewCommentController() *CommentController {
	return &CommentController{}
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
func (commentController *CommentController) GetAllComments(typeResponse string) echo.HandlerFunc  {
	return func(context echo.Context) error {
		commentModel := &comments.Comments{}

		result, error := commentModel.GetAll()

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
GetCommentByID - Get Commetn By ID Controller Method
 */
// @Tags Comments
// @Summary Comment Record By ID godoc
// @Summary Show Comment Record By ID
// @Description get  Comment record By ID
// @Accept  json
// @Produce  json
// @Produce  xml
// @Router /json/comments/{id} [get]
// @Router /xml/comments/{id} [get]
// @Param id query string false "id"
// @Success 200 {array} comments.Comments
// @Failure 500
func (commentController *CommentController) GetCommentByID(typeResponse string) echo.HandlerFunc  {
	return func(context echo.Context) error {
		commentModel := &comments.Comments{}
		id := context.Param("id")

		result, err := commentModel.GetByID(id)

		if err != nil{
			panic(err)
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
	commentModel := &comments.Comments{}
	data := context.Request()

	_, error := commentModel.Create(data)

	if error != nil {
		return error
	}
	return context.String(http.StatusOK,"OK")
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
func (commentController *CommentController) DeleteByID(context echo.Context) error  {
	commentModel := &comments.Comments{}
	id := context.QueryParam("id")

	_, err := commentModel.DeleteByID(id)

	if err != nil{
		panic(err)
	}

	return context.String(http.StatusOK,"OK")
}