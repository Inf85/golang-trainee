package conrtollers

import (
	context2 "context"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)
import "../models/posts"
import "../helpers/response"

/*
PostController - Post Controller
*/
type PostController struct {
	service posts.Storage
}

/*
NewPostController - Constructor
*/
func NewPostController(service posts.Storage) *PostController {
	return &PostController{service: service}
}

var (
	MockDB = map[string]*posts.Posts{
		"testPost": &posts.Posts{100, 500, "testPost", "test", []posts.Comments{}},
	}

	postJson = `{"id":500,user_id":100, "title":"testPost", "body":"test","comment":""}`
)

/*
GetAllPosts - Get All Posts From DataBase
*/
// @Tags Posts
// @Summary ListPosts godoc
// @Summary Show all Post in json
// @Description get all Posts records
// @Accept  application/json
// @Accept application/xml
// @Produce  json
// @Produce  xml
// @Router /posts [get]
// @Success 200 {array} posts.Posts
// @Failure 500
func (postController *PostController) GetAllPosts(context echo.Context) error {
	typeResponse := context.Request().Header.Get("Accept")
	result, error := postController.service.GetAll(context2.Background(), []posts.Posts{})

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

/*
GetByID - Get Post Record By ID
*/
// @Tags Posts
// @Summary Post Record By ID godoc
// @Summary Show Post Record By ID
// @Description get  Post record By ID
// @Accept  application/json
// @Accept application/xml
// @Produce  json
// @Produce  xml
// @Router /posts/{id} [get]
// @Param id query string false "id"
// @Success 200 {array} posts.Posts
// @Failure 500
func (postController *PostController) GetByID(context echo.Context) error {
	typeResponse := context.Request().Header.Get("Accept")
	id := context.Param("id")
	result, error := postController.service.GetByID(context2.Background(), []posts.Posts{}, id)

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
	data := context.Request()
	userID, _ := strconv.Atoi(data.FormValue("userId"))
	title := data.FormValue("title")
	body := data.FormValue("body")

	insertData := posts.Posts{
		UserID: userID,
		Title:  title,
		Body:   body,
	}

	postController.service.CreatePost(context2.Background(), &insertData)
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
	id := context.QueryParam("id")

	_, error := postController.service.DeleteByID(context2.Background(), posts.Posts{}, id)

	if error != nil {
		return error
	}
	return context.String(http.StatusOK, "OK")
}
