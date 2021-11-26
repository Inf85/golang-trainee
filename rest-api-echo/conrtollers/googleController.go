package conrtollers

import (
	"github.com/labstack/echo"
	"net/http"
)
import "../service"
import "../models/user"

/*
GoogleController - struct
 */
type GoogleController struct {

}

/*
NewGoogleController - constructor
 */
func NewGoogleController() *GoogleController  {
	return &GoogleController{}
}


/*
GoogleOAuthCallBack - Google Auth Callback Controller
 */
func (g *GoogleController) GoogleOAuthCallBack(context echo.Context) error {
	userModel := user.User{}

	googleService := service.NewGoogleAuth()


	data, err := googleService.GetUserDataFromGoogle(context.Request().FormValue("code"))

	userModel.GetUserByLogin(data.Email)

	if err != nil{
	return context.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return nil
}

/*
Login - Google Auth Login Method
 */
func (g *GoogleController) Login(context echo.Context) error  {
	googleService := service.NewGoogleAuth()

	oauthState :=  googleService.GenerateStateOauthCookie(context)

	url := service.GoogleConfig.AuthCodeURL(oauthState)

	return context.Redirect(http.StatusTemporaryRedirect, url)

	return nil
}
