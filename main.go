package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ujwaldhakal/go-blog-example/auth"
	docs "github.com/ujwaldhakal/go-blog-example/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	auth.RegisterAuthRoutes(r)
	r.Run(":8081")
}
