package post_service

import (
	"github.com/gin-gonic/gin"
	"github.com/ujwaldhakal/go-blog-example/common"
	post_entity "github.com/ujwaldhakal/go-blog-example/post/entity"
)

type postCreateRequest struct {
	Title       string `json:"title"  binding:"required"`
	Description string `json:"description"  binding:"required"`
}

// @BasePath /v1
// Post
// @Summary Creates post
// @Schemes
// @Accept json
// @Produce json
// @Param Body body postCreateRequest true "Parameters should not be empty"
// @Success 200 {object} common.Response
// @Router /posts [post]
func Create(c *gin.Context) {
	var requestBody postCreateRequest

	if err := c.BindJSON(&requestBody); err != nil {
		common.RespondBadRequest(c, err.Error())
		return
	}

	title := requestBody.Title
	description := requestBody.Description
	post := post_entity.Post{Title: title, Description: description}
	post.Create()
	common.RespondCreated(c, common.Response{Message: "Post has been successfully created"})
}
