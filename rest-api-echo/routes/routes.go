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
	postController := conrtollers.NewPostController()
	commentController := conrtollers.NewCommentController()
	googleController := conrtollers.NewGoogleController()
	authController := conrtollers.AuthController{}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/register", authController.SignUp)
	e.POST("/sign-in", authController.SignIn)

	var apiGroup = e.Group("/api")
	{
		apiGroup.Use(UserIdenty)

		/*        Post Routing  */
		apiGroup.GET("/json/posts", postController.GetAllPosts("json"))
		apiGroup.GET("/xml/posts", postController.GetAllPosts("xml"))
		apiGroup.GET("/json/post/:id", postController.GetByID("json", nil))
		apiGroup.GET("/xml/post/:id", postController.GetByID("xml", nil))

		apiGroup.POST("/post", postController.Create)
		apiGroup.DELETE("/post", postController.Delete)

		/* Comments Routing     */

		apiGroup.GET("/json/comments", commentController.GetAllComments("json"))
		apiGroup.GET("/json/comments/:id", commentController.GetCommentByID("json"))
		apiGroup.GET("/xml/comments", commentController.GetAllComments("xml"))
		apiGroup.GET("/xml/comments/:id", commentController.GetCommentByID("xml"))

		apiGroup.POST("/comment", commentController.Create)
		apiGroup.DELETE("/comment", commentController.DeleteByID)
	}

	var googleURL = e.Group("/auth/google")
	{
		googleURL.Any("/callback", googleController.GoogleOAuthCallBack)
		googleURL.Any("/login", googleController.Login)
	}
}
