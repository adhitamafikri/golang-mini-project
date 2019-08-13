package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine, handler *controller.AuthController) {
	group := router.Group("v1/auth")
	group.GET("", handler.Index)
	group.GET("/login", handler.GetLogin)
	group.GET("/register", handler.GetRegister)
}
