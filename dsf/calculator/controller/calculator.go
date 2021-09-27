package controller

import (
	"middleware-mmksi/dsf/calculator/service"
	"middleware-mmksi/dsf/calculator/service/request"
	"middleware-mmksi/server/cors"
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

// Calculator godoc
// @Tags Calculator
// @Summary Get Pertenor
// @Description Get Calculator Pertenor
// @Produce json
// @Param Auth header string true "Auth"
// @Param applicationName header string true "Application Name"
// @Param requestbody body request.TenorRequest true "Request Body"
// @Success 200 {object} response.TenorResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/calculator/perTenor [post]
func (c *dsfPaymentController) GetTenor(gc *gin.Context) {
	cors.AllowCors(gc)
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

// Calculator godoc
// @Tags Calculator
// @Summary Get All Tenor
// @Description Get Calculator All Tenor
// @Produce json
// @Param Auth header string true "Auth"
// @Param applicationName header string true "Application Name"
// @Param requestbody body request.TenorRequest true "Request Body"
// @Success 200 {object} response.TenorResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/calculator/allTenors [post]
func (c *dsfPaymentController) GetAllTenor(gc *gin.Context) {
	cors.AllowCors(gc)
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
