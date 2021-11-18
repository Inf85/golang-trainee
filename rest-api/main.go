package main

import (
	"fmt"
	"net/http"
)
import "./routes"

func main() {

	routes := routes.RegisterRoutes();
	fmt.Println("Server is starting and listening ...")
	http.ListenAndServe(":8080", routes)
}
