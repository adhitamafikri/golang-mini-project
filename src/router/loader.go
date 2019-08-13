package router

import (
	"../controller"
	"../service"
	"github.com/gin-gonic/gin"
)

func LoadRouter(router *gin.Engine) {
	UserRouter(router, &controller.UserController{
		UserService: service.UserServiceHandler(),
	})
	AuthRouter(router, &controller.AuthController{
		AuthService: service.AuthServiceHandler(),
	})
	PostsRouter(router, &controller.PostsController{
		PostsService: service.PostsServiceHandler(),
	})
}
