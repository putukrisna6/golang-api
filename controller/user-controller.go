package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/helper"
	"github.com/putukrisna6/golang-api/service"
)

type UserController interface {
	Update(context *gin.Context)
	Get(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (controller *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildErrorResponse("authorization error", errToken.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusForbidden, res)
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	userUpdateDTO.ID = id
	user := controller.userService.Update(userUpdateDTO)
	res := helper.BuildValidResponse("OK!", user)
	context.JSON(http.StatusOK, res)
}

func (controller *userController) Get(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildErrorResponse("authorization error", errToken.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusForbidden, res)
	}

	claims := token.Claims.(jwt.MapClaims)
	user := controller.userService.Get(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildValidResponse("OK", user)
	context.JSON(http.StatusOK, res)
}
