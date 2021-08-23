package controller

import (
	"middleware-mmksi/dsf/dsf-program/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dsfProgramController struct {
	dsfProgramService service.DsfProgramService
}

type DsfProgramController interface {
	GetAdditionalInsurance(context *gin.Context)
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
