package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "logged in",
	})
}

func (c *authController) Register(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "registered",
	})
}
