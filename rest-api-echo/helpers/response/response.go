package responsehelper

import (
	"github.com/labstack/echo"
)

/*
Error - Error Message Struct
 */
type Error struct {
	Result string `json:"result"`
	Message string `json:"message"`
}

/*
Success - Success Messages Struct
 */
type Success struct {
	Result string `json:"result"`
	Message map[string] interface{} `json:"message"`
}

/*
NewErrorResponse - Response Error
 */
func NewErrorResponse(context echo.Context, statusCode int, message string) *echo.HTTPError {
	context.Logger().Errorf(message)

	return echo.NewHTTPError(statusCode,Error{"Error",message})
}

/*
BadDataErrorResponse - Bad Data Error
 */
func BadDataErrorResponse(statusCode int, message string) *echo.HTTPError {

	return echo.NewHTTPError(statusCode,Error{"Error", message})
}

/*
SuccessDataResponse - Success message response
 */
func SuccessDataResponse(context echo.Context, statusCode int, message map[string]interface{} ) error {

	return context.JSONPretty(200, Success{"Success",message}, " ")
}