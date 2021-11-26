package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)


const oauthGoogleAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

/*
GoogleAuth - Struct
 */
type GoogleAuth struct {

}

/*
NewGoogleAuth - Constructor
 */
func NewGoogleAuth() *GoogleAuth  {

	return &GoogleAuth{}
}

/*
GoogleConfig - Google Oauth2 Config
 */
var GoogleConfig = &oauth2.Config{
	RedirectURL:  "http://8d01-80-93-120-226.ngrok.io/auth/google/callback",
	ClientID:     "695405174969-2u6rldvppke2oephpjpjcif8ekirvi7s.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

/*
GoogleUser - Google response User Struct
 */
type GoogleUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
}

/*
GetUserDataFromGoogle - Retrive Google Data
 */
func (g *GoogleAuth)GetUserDataFromGoogle(code string) (*GoogleUser, error) {
	var u *GoogleUser
	token, err := GoogleConfig.Exchange(context.Background(), code)

	if err != nil{
		return u, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleAPI + token.AccessToken)

	if err != nil {
		return u, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return u, fmt.Errorf("failed read response: %s", err.Error())
	}

	errGoogleUser := json.Unmarshal(contents, &u)
	if errGoogleUser != nil{
		panic(errGoogleUser)
	}

	return u, nil


}

/*
GenerateStateOauthCookie - Google State Cookie
 */
func (g * GoogleAuth)GenerateStateOauthCookie(ctx echo.Context) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	ctx.SetCookie(&cookie)

	return state

}