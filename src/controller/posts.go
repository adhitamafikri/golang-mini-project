package controller

import (
	"net/http"

	"../service"
	"github.com/gin-gonic/gin"
)

type PostsController struct {
	PostsService service.PostsService
}

func (handler *PostsController) Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome to PostsController!",
	})
}

func (handler *PostsController) GetUserPosts(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Getting all user's posts",
	})
}
