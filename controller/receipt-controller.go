package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/putukrisna6/golang-api/cache"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/helper"
	"github.com/putukrisna6/golang-api/service"
)

type ReceiptController interface {
	All(context *gin.Context)
	Show(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	RefreshCache(keys ...string)
}

type receiptController struct {
	receiptService service.ReceiptService
	receiptCache   cache.ReceiptCache
}

func NewReceiptController(receiptService service.ReceiptService, receiptCache cache.ReceiptCache) ReceiptController {
	return &receiptController{
		receiptService: receiptService,
		receiptCache:   receiptCache,
	}
}

func (c *receiptController) All(context *gin.Context) {
	var receipts []entity.Receipt = c.receiptCache.Get("all")
	if receipts == nil {
		receipts = c.receiptService.All()
		c.receiptCache.Set("all", receipts)
	}

	res := helper.BuildValidResponse("OK", receipts)
	context.JSON(http.StatusOK, res)
}

func (c *receiptController) Show(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var arr []entity.Receipt = c.receiptCache.Get(strconv.FormatUint(id, 10))
	if arr == nil {
		var receipt entity.Receipt = c.receiptService.Show(id)
		if (receipt == entity.Receipt{}) {
			res := helper.BuildErrorResponse("failed to retrieve Receipt", "no data with given receiptID", helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}
		arr = append(arr, receipt)
		c.receiptCache.Set(strconv.FormatUint(id, 10), arr)
	}

	res := helper.BuildValidResponse("OK", arr[0])
	context.JSON(http.StatusOK, res)
}

func (c *receiptController) Insert(context *gin.Context) {
	var receiptCreateDTO dto.ReceiptCreateDTO
	errDTO := context.ShouldBind(&receiptCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.receiptService.Insert(receiptCreateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusCreated, response)

	c.RefreshCache("all")
}

func (c *receiptController) Update(context *gin.Context) {
	var receiptUpdateDTO dto.ReceiptUpdateDTO
	errDTO := context.ShouldBind(&receiptUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.receiptService.Update(receiptUpdateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusOK, response)

	c.RefreshCache("all", strconv.FormatUint(receiptUpdateDTO.ID, 10))
}

func (c *receiptController) Delete(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var receipt entity.Receipt = c.receiptService.Show(id)
	if (receipt == entity.Receipt{}) {
		res := helper.BuildErrorResponse("failed to retrieve Receipt", "no data with given receiptID", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	c.receiptService.Delete(receipt)
	message := fmt.Sprintf("Receipt with ID %v successfuly deleted", receipt.ID)
	res := helper.BuildValidResponse(message, helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	c.RefreshCache("all", strconv.FormatUint(id, 10))
}

func (c *receiptController) RefreshCache(keys ...string) {
	for _, key := range keys {
		c.receiptCache.Del(key)
	}
}
