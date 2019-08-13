package router

import (
	controllers "../controller"
	"../service"
	"github.com/gin-gonic/gin"
)

func LoadRouter(router *gin.Engine) {
	UserRouter(router, &controllers.UserController{
		UserService: services.UserServiceHandler(),
	})
}
