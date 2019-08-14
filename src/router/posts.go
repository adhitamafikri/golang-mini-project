package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func PostsRouter(router *gin.Engine, handler *controller.PostsController) {
	group := router.Group("v1/posts")
	group.GET("", handler.Index)
	group.GET("/all", handler.GetAllPosts)
	// group.GET("/user/:user_id", handler.GetUserPosts)
}
