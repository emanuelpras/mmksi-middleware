package controller

import (
	"middleware-mmksi/dsf/payment/service"
	"middleware-mmksi/dsf/payment/service/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dsfProgramController struct {
	dsfProgramService service.DsfProgramService
}

type DsfProgramController interface {
	GetAdditionalInsurance(context *gin.Context)
	GetPackageNames(context *gin.Context)
	GetCarConditions(context *gin.Context)
	GetPackages(context *gin.Context)
	GetUnitByModels(context *gin.Context)
	GetPaymentTypes(context *gin.Context)
	GetVehicleCategory(context *gin.Context)
	GetBranchID(context *gin.Context)
}

func NewDsfProgramController(
	dsfProgramService service.DsfProgramService,
) *dsfProgramController {
	return &dsfProgramController{
		dsfProgramService: dsfProgramService,
	}
}

func (c *dsfProgramController) GetAdditionalInsurance(gc *gin.Context) {

	res, err := c.dsfProgramService.GetAdditionalInsurance()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetPackageNames(gc *gin.Context) {

	res, err := c.dsfProgramService.GetPackageNames()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetCarConditions(gc *gin.Context) {

	res, err := c.dsfProgramService.GetCarConditions()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetPackages(gc *gin.Context) {
	var applicationName request.HeaderPackageRequest
	if err := gc.ShouldBindHeader(&applicationName); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var packageRequest request.PackageRequest
	if err := gc.ShouldBindJSON(&packageRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetPackages(applicationName, packageRequest)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetUnitByModels(gc *gin.Context) {
	var applicationName request.HeaderUnitByModelsRequest
	if err := gc.ShouldBindHeader(&applicationName); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetUnitByModels(applicationName)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetPaymentTypes(gc *gin.Context) {

	res, err := c.dsfProgramService.GetPaymentTypes()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetVehicleCategory(gc *gin.Context) {
	res, err := c.dsfProgramService.GetVehicleCategory()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetBranchID(gc *gin.Context) {

	res, err := c.dsfProgramService.GetBranchID()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
