package post_service_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	auth_middleware "github.com/ujwaldhakal/go-blog-example/auth/middleware"
	"github.com/ujwaldhakal/go-blog-example/auth/repository"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	db2 "github.com/ujwaldhakal/go-blog-example/db"
	post_entity "github.com/ujwaldhakal/go-blog-example/post/entity"
	post_service "github.com/ujwaldhakal/go-blog-example/post/service"
	"github.com/ujwaldhakal/go-blog-example/user"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupRouter() (*gin.Engine, *httptest.ResponseRecorder) {
	r := gin.Default()

	db := db2.GetConnection()
	db.AutoMigrate(&post_entity.Post{})

	r.Use(auth_middleware.AuthJwtHeaderToken)
	{
		r.POST("v1/posts", post_service.Create)
	}

	w := httptest.NewRecorder()

	return r, w
}

func TestPostCreationWhenEmptyPayloadIsProvided(t *testing.T) {
	r, w := setupRouter()
	defer TearDown()
	req, _ := http.NewRequest("POST", "/v1/posts", nil)
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}

	fmt.Println(obj)
	assert.Equal(t, "invalid request", obj["message"])
}

func TestPostCreationWhenPayloadIsProvided(t *testing.T) {
	r, w := setupRouter()
	defer TearDown()
	var jsonStr = []byte(`{"title":"blog1", "description": "description here"}`)
	req, _ := http.NewRequest("POST", "/v1/posts", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+getUser1Token())
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}

	assert.Equal(t, "Post has been successfully created", obj["message"])
	assert.Equal(t, "success", obj["status"])
}

func getUser1Token() string {
	db := db2.GetConnection()
	password, _ := repository.HashPassword("password")
	email := "john@doe.com"
	db.Create(&user.User{
		ID:          0,
		Name:        "ujwal",
		Email:       email,
		Password:    password,
		Birthday:    nil,
		ActivatedAt: sql.NullTime{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	})

	token, _ := auth.GenerateJwtToken(email)
	return token
}
func TearDown() {
	db := db2.GetConnection()
	db.Exec("Truncate table posts")
}
