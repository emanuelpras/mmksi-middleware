package controller

import (
	"middleware-mmksi/dsf/calculator/service"
	"middleware-mmksi/dsf/calculator/service/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dsfPaymentController struct {
	dsfPaymentService service.DsfPaymentService
}

type DsfPaymentController interface {
	GetTenor(context *gin.Context)
	GetAllTenor(context *gin.Context)
}

func NewDsfPaymentController(
	dsfPaymentService service.DsfPaymentService,
) *dsfPaymentController {
	return &dsfPaymentController{
		dsfPaymentService: dsfPaymentService,
	}
}

func (c *dsfPaymentController) GetTenor(gc *gin.Context) {
	var params request.HeaderTenorRequest
	if err := gc.ShouldBindHeader(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reqBody request.TenorRequest
	if err := gc.ShouldBind(&reqBody); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfPaymentService.GetTenor(params, reqBody)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfPaymentController) GetAllTenor(gc *gin.Context) {
	var params request.HeaderTenorRequest
	if err := gc.ShouldBindHeader(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reqBody request.TenorRequest
	if err := gc.ShouldBind(&reqBody); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfPaymentService.GetAllTenor(params, reqBody)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
