package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/middleware-poc/service"
	"github.com/refactory-id/middleware-poc/service/mrp"
)

type mrpController struct {
	mrpService service.MrpService
}

type MrpController interface {
	GetVehicles(context *gin.Context)
	GetRegions(context *gin.Context)
	GetPrediction(context *gin.Context)
}

func NewMrpController(
	mrpService service.MrpService,
) *mrpController {
	return &mrpController{
		mrpService: mrpService,
	}
}

func (c *mrpController) GetVehicles(gc *gin.Context) {
	var form mrp.GetVehicleRequest
	if err := gc.ShouldBindQuery(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mrpService.GetVehicles(form)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *mrpController) GetRegions(gc *gin.Context) {
	var form mrp.GetRegionsRequest
	if err := gc.ShouldBindQuery(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mrpService.GetRegions(form)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *mrpController) GetPrediction(gc *gin.Context) {
	var form mrp.PredictionRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mrpService.GetPrediction(form)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
