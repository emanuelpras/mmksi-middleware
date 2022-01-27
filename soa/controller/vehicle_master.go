package controller

import (
	"middleware-mmksi/server/cors"
	"middleware-mmksi/soa/service"
	"middleware-mmksi/soa/service/request"
	"net/http"

	er "middleware-mmksi/soa/response"

	"github.com/gin-gonic/gin"
)

type soaControlller struct {
	soaService service.SoaService
}

type SoaController interface {
	VehicleMasterList(context *gin.Context)
	VehicleMasterByAssetCode(context *gin.Context)
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

func (c *soaControlller) VehicleMasterByAssetCode(gc *gin.Context) {
	cors.AllowCors(gc)

	assetCode := gc.Param("assetCode")

	if assetCode == " " {
		gc.JSON(http.StatusBadRequest, &er.ErrorResponse{
			ErrorID: 400,
			Msg: map[string]string{
				"en": "Asset code cannot be empty",
				"id": "Asset code harus diisi",
			},
		})
		return
	}

	result, err := c.soaService.VehicleMasterByAssetCode(assetCode)

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}
