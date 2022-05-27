package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  string      `json:"status" format:"string"`
	Message string      `json:"message" format:"string"`
	Code    int         `json:"code" format:"int"`
	Data    interface{} `json:"code" format:"json"`
}

func respondWithSuccess(response *Response) gin.H {
	g := gin.H{"message": response.Message}
	g["status"] = "success"
	if response.Data != nil {
		g["data"] = response.Data
	}

	return g
}

func respondWithError(response *Response) gin.H {
	g := gin.H{"message": response.Message}
	g["status"] = "error"
	if response.Data != nil {
		g["data"] = response.Data
	}

	return g
}

func RespondBadRequest(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, respondWithError(&Response{Message: errorMessage}))
}

func RespondCreated(c *gin.Context, response Response) {
	c.JSON(http.StatusCreated, respondWithSuccess(&response))
}

func RespondUnauthorized(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnauthorized, respondWithError(&Response{Message: errorMessage}))
}

func RespondOk(c *gin.Context, response Response) {
	c.JSON(http.StatusOK, respondWithSuccess(&response))
}
