package controller

import (
	"net/http"
	"os"

	"middleware-mmksi/salesforce/service"
	"middleware-mmksi/salesforce/service/request"
	"middleware-mmksi/server/cors"

	"github.com/gin-gonic/gin"
)

type salesforceController struct {
	salesforceService service.SalesforceService
}

type SalesforceController interface {
	GetTokenSales(context *gin.Context)
	GetServiceHistory(context *gin.Context)
	GetSparepartSalesHistory(context *gin.Context)
	CheckToken(context *gin.Context)
}

func NewSalesforceController(
	salesforceService service.SalesforceService,
) *salesforceController {
	return &salesforceController{
		salesforceService: salesforceService,
	}
}

var SalesToken = request.SalesRequestAuthorization{}

func (c *salesforceController) GetTokenSales(gc *gin.Context) {

	res, err := c.salesforceService.GetTokenSales()
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	SalesToken = request.SalesRequestAuthorization{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
		InstanceURL: res.InstanceURL,
	}

	if res.AccessToken != "" {
		gc.Next()
	} else {
		gc.Abort()
	}
}

// Salesforce godoc
// @Tags Salesforce Service History
// @Summary Get Service History
// @Description Get Service History from Salesforce
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.ServiceHistoryRequest true "Service History"
// @Success 200 {object} response.ServiceHistoryResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /salesforce/services/serviceHistory [post]
func (c *salesforceController) GetServiceHistory(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.ServiceHistoryRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.salesforceService.GetServiceHistory(form, SalesToken)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Salesforce godoc
// @Tags Salesforce Sparepart Sales History
// @Summary Get Sparepart Sales History
// @Description Get Sparepart Sales History from Salesforce
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.SparepartSalesHistoryRequest true "Sparepart Sales History"
// @Success 200 {object} response.ServiceHistoryResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /salesforce/services/sparepartSalesHistory [post]
func (c *salesforceController) GetSparepartSalesHistory(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.SparepartSalesHistoryRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.salesforceService.GetSparepartSalesHistory(form, SalesToken)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, res)
}

func (c *salesforceController) CheckToken(gc *gin.Context) {

	if SalesToken.AccessToken != "" && os.Getenv("STATUSCODE") != "401" {
		return
	} else if SalesToken.AccessToken != "" && os.Getenv("STATUSCODE") == "401" || SalesToken.AccessToken == "" {
		c.GetTokenSales(gc)
		return
	}
}
