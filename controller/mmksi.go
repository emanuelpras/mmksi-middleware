package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/middleware-poc/service"
	"github.com/refactory-id/middleware-poc/service/mmksi"
)

type mmksiController struct {
	mmksiService service.MmksiService
}

type MmksiController interface {
	GetToken(context *gin.Context)
	GetVehicle(context *gin.Context)
}

func NewMmksiController(
	mmksiService service.MmksiService,
) *mmksiController {
	return &mmksiController{
		mmksiService: mmksiService,
	}
}

var DnetToken = mmksi.VehicleRequestAuthorization{}

func (c *mmksiController) GetToken(gc *gin.Context) {

	res, err := c.mmksiService.GetToken(mmksi.TokenRequest{})
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DnetToken = mmksi.VehicleRequestAuthorization{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
	}

	if res.AccessToken != "" {
		gc.Next()
	} else {
		gc.Abort()
	}
}

func (c *mmksiController) GetVehicle(gc *gin.Context) {
	var form mmksi.VehicleRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}

	res, err := c.mmksiService.GetVehicle(form, DnetToken)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
