package auth_test

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() (*gin.Engine, *httptest.ResponseRecorder) {
	r := gin.Default()

	r.POST("v1/login",auth.Login)

	w := httptest.NewRecorder()

	return r,w
}

func TestLoginWhenEmptyPayloadIsProvided(t *testing.T) {

	r,w := setupRouter()
	req, _ := http.NewRequest("POST", "/v1/login", nil)
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}

	assert.Equal(t,"invalid request",obj["error"])
}


func TestLoginWhenInvalidPayloadIsProvided(t *testing.T) {

	r,w := setupRouter()
	var jsonStr = []byte(`{"username":"username", "password": "password"}`)

	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}

	assert.Equal(t,"Sorry username is incorrect",obj["message"])
}
