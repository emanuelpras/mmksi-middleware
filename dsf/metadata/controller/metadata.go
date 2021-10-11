package controller

import (
	"middleware-mmksi/dsf/metadata/service"
	"middleware-mmksi/dsf/metadata/service/request"
	"middleware-mmksi/server/cors"
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
	GetVariants(context *gin.Context)
	GetPaymentTypes(context *gin.Context)
	GetBrands(context *gin.Context)
	GetModels(context *gin.Context)
	GetVehicleCategory(context *gin.Context)
	GetBranchID(context *gin.Context)
	GetInsuranceTypes(context *gin.Context)
	GetInsurance(context *gin.Context)
	GetAssetCode(context *gin.Context)
	GetProvinces(context *gin.Context)
	GetCities(context *gin.Context)
}

func NewDsfProgramController(
	dsfProgramService service.DsfProgramService,
) *dsfProgramController {
	return &dsfProgramController{
		dsfProgramService: dsfProgramService,
	}
}

// Additional Insurance godoc
// @Tags Metadata
// @Summary Get Additional Insurance
// @Description Get Additional Insurance
// @Produce json
// @Param Authentication header string true "Authentication"
// @Success 200 {object} response.AdditionalInsuranceResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/additionalInsurance [get]
func (c *dsfProgramController) GetAdditionalInsurance(gc *gin.Context) {

	cors.AllowCors(gc)
	res, err := c.dsfProgramService.GetAdditionalInsurance()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Package Names godoc
// @Tags Metadata
// @Summary Get Package Names
// @Description Get Package Names
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param applicationName header string true "Application Name"
// @Param assetCode header string true "Asset Code"
// @Param branchCode header string true "Branch Code"
// @Param carCondition query request.ParamsPackageNameRequest true "Car Condition"
// @Success 200 {object} response.PackageNameResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/packageNames [get]
func (c *dsfProgramController) GetPackageNames(gc *gin.Context) {

	cors.AllowCors(gc)
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

// Car Condition godoc
// @Tags Metadata
// @Summary Get Car Condition
// @Description Get Car Condition
// @Produce json
// @Param Authentication header string true "Authentication"
// @Success 200 {object} response.CarConditionResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/carConditions [get]
func (c *dsfProgramController) GetCarConditions(gc *gin.Context) {

	cors.AllowCors(gc)
	res, err := c.dsfProgramService.GetCarConditions()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Packages godoc
// @Tags Metadata
// @Summary Get Packages
// @Description Get Packages
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param applicationName header string true "Application Name"
// @Param requestbody body request.PackageRequest true "Package"
// @Success 200 {object} response.PackageResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/packages [post]
func (c *dsfProgramController) GetPackages(gc *gin.Context) {

	cors.AllowCors(gc)
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

// Variants godoc
// @Tags Metadata
// @Summary Get Variants
// @Description Get Variants
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param applicationName header string true "Application Name"
// @Success 200 {object} response.VariantsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/variants [get]
func (c *dsfProgramController) GetVariants(gc *gin.Context) {

	cors.AllowCors(gc)
	var applicationName request.HeaderUnitByModelsRequest
	if err := gc.ShouldBindHeader(&applicationName); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetVariants(applicationName)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Payment Types godoc
// @Tags Metadata
// @Summary Get Payment Types
// @Description Get Payment Types
// @Produce json
// @Param Authentication header string true "Authentication"
// @Success 200 {object} response.PaymentTypesResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/paymentTypes [get]
func (c *dsfProgramController) GetPaymentTypes(gc *gin.Context) {

	cors.AllowCors(gc)
	res, err := c.dsfProgramService.GetPaymentTypes()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Models godoc
// @Tags Metadata
// @Summary Get Models
// @Description Get Models
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param brands query request.ModelsRequest true "Models"
// @Success 200 {object} response.ModelsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/models [get]
func (c *dsfProgramController) GetModels(gc *gin.Context) {

	cors.AllowCors(gc)
	var params request.ModelsRequest
	if err := gc.ShouldBindQuery(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.dsfProgramService.GetModels(params)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	gc.JSON(http.StatusOK, res)
}

// Brands godoc
// @Tags Metadata
// @Summary Get Brands
// @Description Get Brands
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param brands query request.BrandsRequest true "Brands Request"
// @Success 200 {object} response.BrandsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/brands [get]
func (c *dsfProgramController) GetBrands(gc *gin.Context) {

	cors.AllowCors(gc)
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

// Vehicle Category godoc
// @Tags Metadata
// @Summary Get Vehicle Category
// @Description Get Vehicle Category
// @Produce json
// @Param Authentication header string true "Authentication"
// @Success 200 {object} response.VehicleCategory
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/vehicleCategory [get]
func (c *dsfProgramController) GetVehicleCategory(gc *gin.Context) {

	cors.AllowCors(gc)
	res, err := c.dsfProgramService.GetVehicleCategory()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// BranchID godoc
// @Tags Metadata
// @Summary Get BranchID
// @Description Get BranchID
// @Produce json
// @Param Authentication header string true "Authentication"
// @Success 200 {object} response.BranchResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/branchID [get]
func (c *dsfProgramController) GetBranchID(gc *gin.Context) {

	cors.AllowCors(gc)
	res, err := c.dsfProgramService.GetBranchID()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Insurance Types godoc
// @Tags Metadata
// @Summary Get Insurance Types
// @Description Get Insurance Types
// @Produce json
// @Param Authentication header string true "Authentication"
// @Success 200 {object} response.InsuranceTypesResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/insuranceTypes [get]
func (c *dsfProgramController) GetInsuranceTypes(gc *gin.Context) {

	cors.AllowCors(gc)
	res, err := c.dsfProgramService.GetInsuranceTypes()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Insurances godoc
// @Tags Metadata
// @Summary Get Insurances
// @Description Get Insurances
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param dsfBranchId query string true "Dsf Branch ID"
// @Param VehicleCategory query string true "Vehicle Category"
// @Param InsuranceTypeCode query string true "Insurance Type Code"
// @Param Car Condition query string true "Car Condition"
// @Success 200 {object} response.InsuranceResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/insurances [get]
func (c *dsfProgramController) GetInsurance(gc *gin.Context) {

	cors.AllowCors(gc)
	var params request.InsuranceRequest
	if err := gc.ShouldBindQuery(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetInsurance(params)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Asset Code godoc
// @Tags Metadata
// @Summary Get Asset Code
// @Description Get Asset Code
// @Consume json
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param applicationName header string true "Application Name"
// @Param requestbody body request.AssetCodeRequest true "Asset Code Request"
// @Success 200 {object} response.AssetCodeResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/assetCode [post]
func (c *dsfProgramController) GetAssetCode(gc *gin.Context) {

	cors.AllowCors(gc)
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

// Province godoc
// @Tags Metadata
// @Summary Get Province
// @Description Get Province
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param province query request.ProvincesRequest true "Province Request"
// @Success 200 {object} response.ProvincesResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/province [get]
func (c *dsfProgramController) GetProvinces(gc *gin.Context) {

	cors.AllowCors(gc)
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

// Cities godoc
// @Tags Metadata
// @Summary Get Cities
// @Description Get Cities
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param cities query request.CitiesRequest true "Cities Request"
// @Success 200 {object} response.CitiesResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /dsf/metadata/cities [get]
func (c *dsfProgramController) GetCities(gc *gin.Context) {

	cors.AllowCors(gc)
	var params request.CitiesRequest
	if err := gc.ShouldBindQuery(&params); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.dsfProgramService.GetCities(params)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusOK, res)

}
