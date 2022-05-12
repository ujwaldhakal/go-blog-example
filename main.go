package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ujwaldhakal/go-blog-example/auth"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth.RegisterAuthRoutes(r)
	r.Run(":8081")
}
