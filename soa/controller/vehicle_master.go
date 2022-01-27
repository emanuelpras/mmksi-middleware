package controller

import (
	"middleware-mmksi/server/cors"
	"middleware-mmksi/soa/service"
	"middleware-mmksi/soa/service/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type soaControlller struct {
	soaService service.SoaService
}

type SoaController interface {
	VehicleMasterList(context *gin.Context)
}

func NewSoaController(soaService service.SoaService) *soaControlller {
	return &soaControlller{
		soaService: soaService,
	}
}

func (c *soaControlller) VehicleMasterList(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.SoaVehicleMasterRequest

	if err := gc.ShouldBindQuery(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.soaService.VehicleMasterList(form)

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}
