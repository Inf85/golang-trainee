package routes

import (
	"github.com/gorilla/mux"
)
import "../controllers"

/*
RegisterRoutes - Routes List
 */
func RegisterRoutes() *mux.Router {
	indexController := controllers.IndexController{}
	postController := controllers.PostController{}
	commentController := controllers.CommentController{}

	r := mux.NewRouter()
	r.HandleFunc("/", indexController.Index).Methods("GET")
	r.HandleFunc("/json/posts", postController.GetAllPosts( "json")).Methods("GET")
	r.HandleFunc("/post", postController.CreatePost).Methods("POST")
	r.HandleFunc("/post/{id:[0-9]+}", postController.DeletePostByID).Methods("DELETE")
	r.HandleFunc("/json/posts/{id:[0-9]+}", postController.GetPostByID( "json")).Methods("GET")
	r.HandleFunc("/xml/posts", postController.GetAllPosts( "xml")).Methods("GET")


	r.HandleFunc("/json/comments", commentController.GetAllComments( "json")).Methods("GET")
	return r
}
