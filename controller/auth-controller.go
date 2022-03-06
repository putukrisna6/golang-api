package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/helper"
	"github.com/putukrisna6/golang-api/service"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(context *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := context.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildValidResponse("OK", v)
		context.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("invalid credentials", "invalid credentials", helper.EmptyObj{})
	context.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(context *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := context.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("failed to process request", "duplicate email", helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}
	newUser := c.authService.CreateUser(registerDTO)
	token := c.jwtService.GenerateToken(strconv.FormatUint(newUser.ID, 10))
	newUser.Token = token
	response := helper.BuildValidResponse("OK", newUser)
	context.JSON(http.StatusCreated, response)
}
