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
