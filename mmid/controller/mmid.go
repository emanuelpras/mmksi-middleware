package controller

import (
	"net/http"

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
	var header request.HeaderRequest
	var form request.ServiceHistoryRequest
	if err := gc.ShouldBindHeader(&header); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := gc.ShouldBindJSON(&form); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.mmidService.GetServiceHistory(header, form)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusOK, res)
}
