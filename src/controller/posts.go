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

func (handler *PostsController) GetAllPosts(context *gin.Context) {
	results := handler.PostsService.GetAllPosts()
	context.JSON(http.StatusOK, results)
}
