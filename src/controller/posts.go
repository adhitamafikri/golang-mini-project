package controller

import (
	"net/http"
	"strconv"

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

func (handler *PostsController) GetPostsByUserID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("userID"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	results := handler.PostsService.GetPostsByUserID(id)
	context.JSON(http.StatusOK, results)
}
