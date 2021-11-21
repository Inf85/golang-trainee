package main

import "github.com/labstack/echo"
import "./routes"


func main()  {
	e := echo.New()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
