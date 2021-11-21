package routes

import "github.com/labstack/echo"
import "../conrtollers"

/*
RegisterRoutes - Routes List
 */
func RegisterRoutes(e *echo.Echo)  {
	postController := conrtollers.PostController{}

	e.GET("/json/posts", postController.GetAllPosts("json"))
	e.GET("/xml/posts", postController.GetAllPosts("xml"))
	e.GET("/json/post/:id", postController.GetByID("json"))
	e.GET("/xml/post/:id", postController.GetByID("xml"))
}
