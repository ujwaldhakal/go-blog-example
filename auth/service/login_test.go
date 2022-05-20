package auth_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ujwaldhakal/go-blog-example/auth/repository"
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

	db := db2.GetConnection()
	db.AutoMigrate(&user.User{})

	r.POST("v1/login", auth.Login)

	w := httptest.NewRecorder()

	return r, w
}

func TestLoginWhenEmptyPayloadIsProvided(t *testing.T) {

	r, w := setupRouter()
	defer TearDown()
	req, _ := http.NewRequest("POST", "/v1/login", nil)
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}

	assert.Equal(t, "invalid request", obj["error"])
}

func TestLoginWhenInvalidPayloadIsProvided(t *testing.T) {

	r, w := setupRouter()
	defer TearDown()
	var jsonStr = []byte(`{"email":"username", "password": "password"}`)

	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}

	assert.Equal(t, auth.UsernameOrPasswordIncorrect, obj["message"])
}

func TestLoginWhenValidCredentialsIsProvided(t *testing.T) {
	r, w := setupRouter()
	TearDown()
	hydrateData()
	var jsonStr = []byte(`{"email":"john@doe.com", "password": "password"}`)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}
	token := obj["data"].(map[string]interface{})["token"]
	assert.NotEmpty(t, token)
}

func TearDown() {
	db := db2.GetConnection()
	db.Exec("Truncate table users")
}

func hydrateData() {
	db := db2.GetConnection()
	password, _ := repository.HashPassword("password")
	db.Create(&user.User{
		ID:          0,
		Name:        "ujwal",
		Email:       "john@doe.com",
		Password:    password,
		Birthday:    nil,
		ActivatedAt: sql.NullTime{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	})
}
