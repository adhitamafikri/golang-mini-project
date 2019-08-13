package controller

import (
	"net/http"

	"../service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
}

func (handler *AuthController) Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome to AuthController!",
	})
}

func (handler *AuthController) GetLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Getting Login Screen",
	})
}

func (handler *AuthController) GetRegister(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Getting Register Screen",
	})
}
