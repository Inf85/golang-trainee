package responsehelper

import (
	"github.com/labstack/echo"
)

type Error struct {
	Result string `json:"result"`
	Message string `json:"message"`
}

type Success struct {
	Result string `json:"result"`
	Message map[string] interface{} `json:"message"`
}

func NewErrorResponse(context echo.Context, statusCode int, message string) *echo.HTTPError {
	context.Logger().Errorf(message)

	return echo.NewHTTPError(statusCode,Error{"Error",message})
}

func BadDataErrorResponse(statusCode int, message string) *echo.HTTPError {

	return echo.NewHTTPError(statusCode,Error{"Error", message})
}

func SuccessDataResponse(context echo.Context, statusCode int, message map[string]interface{} ) error {

	return context.JSONPretty(200, Success{"Success",message}, " ")
}