package controllers

import (
	"fmt"
	"net/http"
)
/*
IndexController struct
 */
type IndexController struct{}


/*
Index - home page
 */
func (IndexController) Index(writer http.ResponseWriter, request *http.Request)  {

	fmt.Fprint(writer, "Index Page")
}