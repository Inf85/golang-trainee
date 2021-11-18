package routes

import "github.com/gorilla/mux"
import "../pages"

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", pages.IndexPageHandler).Methods("GET")
	return r
}
