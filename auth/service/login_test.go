package auth_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	db2 "github.com/ujwaldhakal/go-blog-example/db"
	"github.com/ujwaldhakal/go-blog-example/user"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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


func TestLoginWhenValidCredentialsIsProdivded(t *testing.T) {
	r,w := setupRouter()
	hydrateData()
	var jsonStr = []byte(`{"username":"john@doe.com", "password": "password"}`)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}
	token := obj["data"].(map[string]interface{})["token"]
	assert.NotEmpty(t,token)
}

func hydrateData()  {
	db := db2.GetConnection()
	db.AutoMigrate(&user.User{})
	db.Create(&user.User{
		ID:          0,
		Name:        "ujwal",
		Email:       "john@doe.com",
		Password:    "password",
		Birthday:    nil,
		ActivatedAt: sql.NullTime{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	})
}