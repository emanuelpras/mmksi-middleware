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
	GetVehicles(context *gin.Context)
}

func NewMmksiController(
	mmksiService service.MmksiService,
) *mmksiController {
	return &mmksiController{
		mmksiService: mmksiService,
	}
}

func (c *mmksiController) GetToken(gc *gin.Context) {

	var form mmksi.TokenRequest
	if err := gc.ShouldBind(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmksiService.GetToken(form)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *mmksiController) GetVehicles(gc *gin.Context) {

	var form mmksi.VehicleRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var formHeader mmksi.VehicleRequestAuthorization
	if err := gc.ShouldBindHeader(&formHeader); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmksiService.GetVehicles(form, formHeader)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
