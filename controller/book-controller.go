package controller

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/putukrisna6/golang-api/service"
)

type BookController interface {
	All(context *gin.Context)
	Get(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookController(bookService service.BookService, jwtService service.JWTService) BookController {
	return &bookController{
		bookService: bookService,
		jwtService:  jwtService,
	}
}

func (c *bookController) All(context *gin.Context) {

}

func (c *bookController) Get(context *gin.Context) {

}

func (c *bookController) Insert(context *gin.Context) {

}

func (c *bookController) Update(context *gin.Context) {

}

func (c *bookController) Delete(context *gin.Context) {

}

func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
