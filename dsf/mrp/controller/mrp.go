package controller

import (
	"net/http"

	"middleware-mmksi/dsf/mrp/service"
	"middleware-mmksi/dsf/mrp/service/request"

	"github.com/gin-gonic/gin"
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

// Vehicle godoc
// @Summary Get Vehicle
// @Description Get Vehicle
// @Produce json
// @Param Auth header string true "Auth"
// @Param vehicle query request.GetVehicleRequest true "Vehicle Request"
// @Success 200 {object} response.GetVehiclesResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/tradein/vehicles [get]
func (c *mrpController) GetVehicles(gc *gin.Context) {
	var form request.GetVehicleRequest
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

// Regions godoc
// @Summary Get Regions
// @Description Get Regions
// @Produce json
// @Param Auth header string true "Auth"
// @Param province query request.GetRegionsRequest true "Province"
// @Success 200 {object} response.GetRegionsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/tradein/regions [get]
func (c *mrpController) GetRegions(gc *gin.Context) {
	var form request.GetRegionsRequest
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

// Prediction godoc
// @Summary Get Prediction
// @Description Get Prediction
// @Produce json
// @Param Auth header string true "Auth"
// @Param requestbody body request.PredictionRequest true "Request Body"
// @Success 200 {object} response.PredictionResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/tradein/prediction [post]
func (c *mrpController) GetPrediction(gc *gin.Context) {
	var form request.PredictionRequest
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
