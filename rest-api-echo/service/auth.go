package service

import (
	"../models/user"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/jwt-go"
	"github.com/labstack/echo"
	"time"
)

const salt = "3jna9d90r332n"
const tokenTTL = 12 * time.Hour
const signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

/*
AuthService - Service Autg
*/
type AuthService struct {
	userModel user.User
}

/*
NewAuthService - Constructor
*/
func NewAuthService() *AuthService {
	return &AuthService{}
}

/*
CreateUser - AuthService Password Hash
*/
func (s *AuthService) CreateUser(user user.User) (int, *echo.HTTPError, error) {
	user.Password = generateHashPassword(user.Password)

	return s.userModel.CreateUser(user)
}

/*
GenerateToken - JWT Token Generation
*/
func (s *AuthService) GenerateToken(userName string, password string) (string, error) {
	user, err := s.userModel.GetUser(userName, generateHashPassword(password))

	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

/*
ParsingToken - ParsingToken
*/
func (s *AuthService) ParsingToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, result := token.Method.(*jwt.SigningMethodHMAC); !result {
			return nil, errors.New("Invalid sign in method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("Error claims")
	}

	return claims.UserID, nil
}

/*
generateHashPassword - generate Hash password
*/
func generateHashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
