package conrtollers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)
import "../models/user"
import "../helpers/response"
import "../service"

/*
AuthController - Auth Controller
*/
type AuthController struct {
}

/*
signInInput Login Struct
 */
type signInInput struct {
	Username string `json:"username" binding:"required" form:"user_name"`
	Password string `json:"password" binding:"required" form:"password"`
}

/*
SignUp - Controller Register Method
*/
func (authController *AuthController) SignUp(context echo.Context) error {
	var input user.User

	error := context.Bind(&input)

	if error != nil {
		responsehelper.NewErrorResponse(context, http.StatusBadRequest, error.Error())
	}
	authService := service.NewService()

	id, _, duplicateError := authService.CreateUser(input)

	if duplicateError != nil {
		return duplicateError
	}

	if duplicateError == nil {
		responsehelper.SuccessDataResponse(context, http.StatusCreated, map[string]interface{}{"ID": strconv.Itoa(id)})
	}

	return nil
}

/*
SignIn - Login Controller Method
 */
func (authController *AuthController) SignIn(context echo.Context) error {
	var input signInInput
	authService := service.NewService()

	error := context.Bind(&input)

	if error != nil {
		responsehelper.NewErrorResponse(context, http.StatusBadRequest, error.Error())
	}

	token, err := authService.GenerateToken(input.Username, input.Password)

	if err != nil{
		responsehelper.NewErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"token" : token,
	})

	return nil
}
