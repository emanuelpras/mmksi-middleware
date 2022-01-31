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

// Mmid godoc
// @Tags Soa Vehicle
// @Summary Get Vehicle Master List
// @Description Get Vehicle Master from SOA
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param vehicle query request.SoaVehicleMasterRequest true "Soa Vehicle List"
// @Success 200 {object} response.ListVehicleMasterResponse
// @Failure 400 {object} er.ErrorResponse
// @Router /soa/metadata/vehicle [get]
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

// Mmid godoc
// @Tags Soa Vehicle
// @Summary Get Vehicle Master
// @Description Get Vehicle Master from SOA
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param vehicle path request.VehicleMasterByAssetCodeRequest true "Soa Vehicle"
// @Success 200 {object} response.VehicleMasterByAssetCodeResponse
// @Failure 400 {object} er.ErrorResponse
// @Router /soa/metadata/vehicle/{assetCode} [get]
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
