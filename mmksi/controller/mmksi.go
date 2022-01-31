package controller

import (
	_ "middleware-mmksi/mmid/response"
	"middleware-mmksi/mmksi/service"
	"middleware-mmksi/mmksi/service/request"
	"middleware-mmksi/server/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mmksiController struct {
	mmksiService service.MmksiService
}

type MmksiController interface {
	GetToken(context *gin.Context)
	GetVehicle(context *gin.Context)
	GetVehicleColor(context *gin.Context)
}

func NewMmksiController(
	mmksiService service.MmksiService,
) *mmksiController {
	return &mmksiController{
		mmksiService: mmksiService,
	}
}

var DnetToken = request.VehicleRequestAuthorization{}

func (c *mmksiController) GetToken(gc *gin.Context) {

	res, err := c.mmksiService.GetToken(request.TokenRequest{})
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DnetToken = request.VehicleRequestAuthorization{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
	}

	if res.AccessToken != "" {
		gc.Next()
	} else {
		gc.Abort()
	}
}

// Vehicle godoc
// @Tags MMKSI Vehicle
// @Summary Get Vehicle
// @Description Get Vehicle from MMKSI
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.VehicleRequest true "Vehicle"
// @Success 200 {object} response.VehicleResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /mmksi/getData [post]
func (c *mmksiController) GetVehicle(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.VehicleRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmksiService.GetVehicle(form, DnetToken)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Vehicle Color godoc
// @Tags MMKSI Vehicle
// @Summary Get Vehicle Color
// @Description Get Vehicle Color from MMKSI
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.VehicleRequest true "Vehicle Color"
// @Success 200 {object} response.VehicleColorResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /mmksi/vehicle [post]
func (c *mmksiController) GetVehicleColor(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.VehicleRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmksiService.GetVehicleColor(form, DnetToken)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
