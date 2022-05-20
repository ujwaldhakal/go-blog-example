package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

	r.POST("v1/register", Register)

	w := httptest.NewRecorder()

	return r, w
}

func TestRegisterWhenEmptyPayloadIsProvided(t *testing.T) {

	r, w := setupRouter()
	defer TearDown()
	req, _ := http.NewRequest("POST", "/v1/register", nil)
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}
	assert.Equal(t, "invalid request", obj["message"])
	assert.Equal(t, "error", obj["status"])
}

func TestRegisterWhenInCompletePayloadIsProvided(t *testing.T) {

	r, w := setupRouter()
	defer TearDown()
	var jsonStr = []byte(`{"email":"username"}`)
	req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}
	assert.Equal(t, "error", obj["status"])
}

func TestRegisterWhenUniqueUsernameIsProvidedPayloadIsProvided(t *testing.T) {

	r, w := setupRouter()
	defer TearDown()
	var jsonStr = []byte(`{"email":"username","password":"test123","confirmation_password":"test123"}`)
	req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}
	assert.Equal(t, "success", obj["status"])
	assert.Equal(t, "Created successfully", obj["message"])
}

func TestRegisterWhenDuplicateEmailIsProvided(t *testing.T) {

	r, w := setupRouter()
	defer TearDown()
	hydrateData()
	var jsonStr = []byte(`{"email":"john@doe.com","password":"test123","confirmation_password":"test123"}`)
	req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	body := w.Body

	var obj map[string]interface{}
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		panic(err)
	}
	assert.Equal(t, "error", obj["status"])
	assert.Equal(t, emailAlreadyExists, obj["message"])
}

func TearDown() {
	db := db2.GetConnection()
	db.Exec("Truncate table users")
}

func hydrateData() {
	db := db2.GetConnection()
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
