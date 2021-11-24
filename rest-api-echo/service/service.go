package service

import (
	"../models/user"
	"github.com/labstack/echo"
)


type Authorization interface {
	CreateUser(user user.User) (int, *echo.HTTPError,error)
	GenerateToken(userName string, password string) (string, error)
	ParsingToken(token string) (int, error)
}

type Service struct {
	Authorization
}


func NewService() *Service {
	return &Service{
		Authorization:NewAuthService(),
	}
}