package controller

import (
	"net/http"
	"strconv"

	_ "middleware-mmksi/mmid/response"
	"middleware-mmksi/mmid/service"
	"middleware-mmksi/mmid/service/request"
	"middleware-mmksi/server/cors"

	"github.com/gin-gonic/gin"
)

type mmidController struct {
	mmidService service.MmidService
}

type MmidController interface {
	GetServiceHistory(context *gin.Context)
	GetServiceHistoryBatch(context *gin.Context)
	GetSparepartList(context *gin.Context)
}

func NewMmidController(
	mmidService service.MmidService,
) *mmidController {
	return &mmidController{
		mmidService: mmidService,
	}
}

// Mmid godoc
// @Tags Mmid Service History
// @Summary Get Service History
// @Description Get Service History from Mmid
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.ServiceHistoryRequest true "Service History"
// @Success 200 {object} response.ServiceHistoryResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /mmid/services/serviceHistory [post]
func (c *mmidController) GetServiceHistory(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.ServiceHistoryRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmidService.GetServiceHistory(form)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if res.Alerts.Code != "200" {
		code, _ := strconv.Atoi(res.Alerts.Code)
		gc.JSON(code, res)
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Mmid godoc
// @Tags Mmid Service History Batch
// @Summary Get Service History Batch
// @Description Get Service History Batch from Mmid
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.ServiceHistoryRequest true "Service History"
// @Success 200 {object} response.ServiceHistoryBatchResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /mmid/services/serviceHistory [post]
func (c *mmidController) GetServiceHistoryBatch(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.BatchServiceHistoryRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmidService.GetServiceHistoryBatch(form)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusOK, res)
}

// Mmid godoc
// @Tags Mmid Spareparts List
// @Summary Get Spareparts List
// @Description Get Sparepart List from Mmid
// @Produce json
// @Param Authentication header string true "Authentication"
// @Param requestbody body request.SparepartListRequest true "Service History"
// @Success 200 {object} response.SparepartListResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /mmid/services/serviceHistory [post]
func (c *mmidController) GetSparepartList(gc *gin.Context) {
	cors.AllowCors(gc)
	var form request.SparepartListRequest
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmidService.GetSparepartList(form)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusOK, res)
}
