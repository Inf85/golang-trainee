package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)
import "../helpers/response"
import "../service"
const authHeader = "Authorization"
const usrContext = "userId"

/*
HandlerFunc - Struct
 */
type HandlerFunc struct {

}

/*
UserIdenty - User Middleware
 */
func UserIdenty(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header.Get(authHeader)
		authService := service.NewService()
		if header == "" {
			err := responsehelper.NewErrorResponse(context, http.StatusUnauthorized, "Auth Header Is Empty")
			return err
		}

		partsOfHeader := strings.Split(header, " ")

		if len(partsOfHeader) != 2 {
			err := responsehelper.NewErrorResponse(context, http.StatusUnauthorized, "Invalid Auth Header")
			return err
		}

		userID, err := authService.ParsingToken(partsOfHeader[1])
		if err != nil {
			return responsehelper.NewErrorResponse(context, http.StatusUnauthorized, err.Error())
		}

		context.Set(usrContext, userID)
		return next(context)
	}
}
