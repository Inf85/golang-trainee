package routes

import (
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
)
import "../conrtollers"

/*
RegisterRoutes - Routes List
 */
func RegisterRoutes(e *echo.Echo)  {
	postController := conrtollers.PostController{}
	commentController := conrtollers.CommentController{}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	/*        Post Routing  */
	e.GET("/json/posts", postController.GetAllPosts("json"))
	e.GET("/xml/posts", postController.GetAllPosts("xml"))
	e.GET("/json/post/:id", postController.GetByID("json"))
	e.GET("/xml/post/:id", postController.GetByID("xml"))

	e.POST("/post", postController.Create)
	e.DELETE("/post", postController.Delete)

	/* Comments Routing     */

	e.GET("/json/comments", commentController.GetAllComments("json"))
	e.GET("/json/comments/:id", commentController.GetCommentByID("json"))
	e.GET("/xml/comments", commentController.GetAllComments("xml"))
	e.GET("/xml/comments/:id", commentController.GetCommentByID("xml"))

	e.POST("/comment", commentController.Create)
	e.DELETE("/comment", commentController.DeleteByID)
}
