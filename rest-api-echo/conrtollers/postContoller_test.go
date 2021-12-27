package conrtollers

import (
	"../models/posts"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockDB = map[string]*posts.Posts{
		"testPost": &posts.Posts{100, 500, "testPost", "test", []posts.Comments{}},
	}

//	postJson = `{"id":500,user_id":100, "title":"testPost", "body":"test","comment":""}`
)

func TestGetByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/json/post/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("500")

	controller := &PostController{}
	if assert.NoError(t, controller.GetByID("json", c)) {

	}
}
