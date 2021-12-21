package routes

import (
	"../conrtollers"
	"../helpers/service-manager"
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
)

/*
RegisterRoutes - Routes List
*/
func RegisterRoutes(e *echo.Echo, serviceManager *servicemanager.ServiceManager) {

	postController := conrtollers.NewPostController(serviceManager.PostService)
	commentController := conrtollers.NewCommentController(serviceManager.CommentService)
	googleController := conrtollers.NewGoogleController()
	authController := conrtollers.AuthController{}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/register", authController.SignUp)
	e.POST("/sign-in", authController.SignIn)

	var apiGroup = e.Group("/api")
	{
		apiGroup.Use(UserIdenty)

		/*        Post Routing  */
		apiGroup.GET("/posts", postController.GetAllPosts)
		apiGroup.GET("/post/:id", postController.GetByID)
		apiGroup.POST("/post", postController.Create)
		apiGroup.DELETE("/post", postController.Delete)

		/* Comments Routing     */

		apiGroup.GET("/comments", commentController.GetAllComments)
		apiGroup.GET("/comments/:id", commentController.GetCommentByID)
		apiGroup.POST("/comment", commentController.Create)
		apiGroup.DELETE("/comment", commentController.DeleteByID)
	}

	var googleURL = e.Group("/auth/google")
	{
		googleURL.Any("/callback", googleController.GoogleOAuthCallBack)
		googleURL.Any("/login", googleController.Login)
	}
}
