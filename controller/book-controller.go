package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/helper"
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
	var books []entity.Book = c.bookService.All()
	res := helper.BuildValidResponse("OK", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) Get(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var book entity.Book = c.bookService.Get(id)
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("failed to retrieve Book", "no data with given bookID", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildValidResponse("OK", book)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	bookCreateDTO.UserID = convertedUserID
	result := c.bookService.Insert(bookCreateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID != nil {
			panic(errID.Error())
		}

		bookUpdateDTO.UserID = id
		result := c.bookService.Update(bookUpdateDTO)
		response := helper.BuildValidResponse("OK", result)
		context.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("you do not have permission to update this Book", "you are not the owner", helper.EmptyObj{})
	context.AbortWithStatusJSON(http.StatusForbidden, response)
}

func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		message := fmt.Sprintf("Book with ID %v successfuly deleted", book.ID)
		res := helper.BuildValidResponse(message, helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
		return
	}

	response := helper.BuildErrorResponse("you do not have permission to delete this Book", "you are not the owner", helper.EmptyObj{})
	context.AbortWithStatusJSON(http.StatusForbidden, response)
}

func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
