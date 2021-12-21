package service

import (
	"../models/user"
	"github.com/labstack/echo"
)

/*
Service - Struct
*/
type Service struct {
	Authorization
	GoogleAuthorization
}


/*
Authorization - Interface
 */
type Authorization interface {
	CreateUser(user user.User) (int, *echo.HTTPError,error)
	GenerateToken(userName string, password string) (string, error)
	ParsingToken(token string) (int, error)
}

/*
GoogleAuthorization - Interface
 */
type GoogleAuthorization interface {
	GenerateStateOauthCookie(context echo.Context) string
	GetUserDataFromGoogle(code string) (*GoogleUser, error)
}



/*
NewService - Constructor
 */
func NewService() *Service {
	return &Service{
		Authorization:NewAuthService(),
		GoogleAuthorization:NewGoogleAuth(),
	}
}