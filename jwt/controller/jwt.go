package controller

import (
	"net/http"

	"middleware-mmksi/jwt/service"
	"middleware-mmksi/jwt/service/request"

	"github.com/gin-gonic/gin"
)

type jwtController struct {
	jwtService service.JwtService
}

type JwtController interface {
	CreateToken(gc *gin.Context)
	RefreshToken(gc *gin.Context)
	Auth(gc *gin.Context)
}

func NewJwtController(
	jwtService service.JwtService,
) *jwtController {
	return &jwtController{
		jwtService: jwtService,
	}
}

func (c *jwtController) CreateToken(gc *gin.Context) {
	var paramJwt request.TokenMmksiRequest
	if err := gc.ShouldBindHeader(&paramJwt); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.jwtService.CreateToken(gc, paramJwt)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	gc.JSON(http.StatusOK, res)
}

func (c *jwtController) RefreshToken(gc *gin.Context) {
	var paramJwt request.TokenRefreshRequest
	if err := gc.ShouldBindHeader(&paramJwt); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.jwtService.RefreshToken(gc, paramJwt)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	gc.JSON(http.StatusOK, res)
}

func (c *jwtController) Auth(gc *gin.Context) {
	var auth request.AuthRequest
	if err := gc.ShouldBindHeader(&auth); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.jwtService.Auth(gc, auth)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
}
