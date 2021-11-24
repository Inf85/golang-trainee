package user

import (
	"../../database"
	"../../helpers/response"
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
)

/*
User - User Struct
 */
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" binding:"required" form:"login"`
	Password string `json:"password" binding:"required" form:"password"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
}

/*
CreateUser - Create User in DataBase
 */
func (u *User) CreateUser(user User) (int,*echo.HTTPError, error) {
	db := dbconnect.SetDBConnection()
	result := db.Where("login = ?", user.Login).First(&user)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {

		return 0, nil,responsehelper.BadDataErrorResponse(http.StatusBadRequest, "User With Same Login in exists")
	}

	postData := User{Login: user.Login, Password: user.Password, Name: user.Name}

	db.Create(&postData)

	return postData.ID, nil, nil
}

/*
GetUser - Get USer By login and Password
 */
func (u *User) GetUser(userName string, password string) (User, error){
	var user User

	db := dbconnect.SetDBConnection()
	db.Where("login = ? AND password = ?", userName, password).First(&user)

	return user, nil
}
