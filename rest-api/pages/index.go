package pages

import (
	"fmt"
	"net/http"
)

func IndexPageHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Index Page")
}
