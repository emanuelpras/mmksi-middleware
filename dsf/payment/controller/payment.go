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
	GetBrands(context *gin.Context)
	GetVehicleCategory(context *gin.Context)
	GetBranchID(context *gin.Context)
	GetInsuranceTypes(context *gin.Context)
	GetInsurance(context *gin.Context)
	GetAssetCode(context *gin.Context)
	GetProvinces(context *gin.Context)
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
	var params request.HeaderPackageNameRequest
	if err := gc.ShouldBindHeader(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var queryParams request.ParamsPackageNameRequest
	if err := gc.ShouldBindQuery(&queryParams); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetPackageNames(params, queryParams)
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

func (c *dsfProgramController) GetBrands(gc *gin.Context) {
	var brandRequest request.BrandsRequest
	if err := gc.ShouldBindQuery(&brandRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.dsfProgramService.GetBrands(brandRequest)
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

func (c *dsfProgramController) GetInsuranceTypes(gc *gin.Context) {

	res, err := c.dsfProgramService.GetInsuranceTypes()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetInsurance(gc *gin.Context) {
	var params request.InsuranceRequest
	if err := gc.ShouldBindQuery(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetInsurance(params)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetAssetCode(gc *gin.Context) {
	var applicationName request.HeaderAssetCodeRequest
	if err := gc.ShouldBindHeader(&applicationName); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var assetCodeRequest request.AssetCodeRequest
	if err := gc.ShouldBindJSON(&assetCodeRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetAssetCode(applicationName, assetCodeRequest)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *dsfProgramController) GetProvinces(gc *gin.Context) {
	var params request.ProvincesRequest
	if err := gc.ShouldBindQuery(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetProvinces(params)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}
