package main_test

import (
	"github.com/labstack/echo"
	"testing"
)
import "./routes"

func TestMain(m *testing.M)  {

	e := echo.New()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}