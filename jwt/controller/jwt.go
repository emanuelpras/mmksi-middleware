package controller

import (
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/service"
	"middleware-mmksi/jwt/service/request"
	_ "middleware-mmksi/mmid/response"
	"middleware-mmksi/server/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jwtController struct {
	jwtService service.JwtService
}

type JwtController interface {
	CreateToken(gc *gin.Context)
	RefreshToken(gc *gin.Context)
	Auth(gc *gin.Context)
	SigninAws(gc *gin.Context)
	ReSigninAws(gc *gin.Context)
}

func NewJwtController(
	jwtService service.JwtService,
) *jwtController {
	return &jwtController{
		jwtService: jwtService,
	}
}

// Authenticate godoc
// Tags Token
// Summary Provides a JSON Web Token
// Description Authenticates a user and provides a JWT to Authorize API calls
// Consume application/x-www-form-urlencoded
// Produce json
// Param company header string true "Company"
// Success 200 {object} response.TokenMmksiResponse
// Failure 400 {object} response.ErrorResponse
// Router /auth/token [post]
func (c *jwtController) CreateToken(gc *gin.Context) {
	cors.AllowCors(gc)
	var paramJwt request.HeaderTokenRequest
	if err := gc.ShouldBindHeader(&paramJwt); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.jwtService.CreateToken(gc, paramJwt, repo.Timeout{})
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	gc.JSON(http.StatusOK, res)
}

// Authenticate godoc
// Tags Token
// Summary Provides a Refresh Token
// Description Authenticates a user and provides a JWT to Authorize API calls
// Consume application/x-www-form-urlencoded
// Produce json
// Param refreshToken header string true "Refresh Token"
// Success 200 {object} response.TokenMmksiResponse
// Failure 400 {object} response.ErrorResponse
// Router /token/refresh [post]
func (c *jwtController) RefreshToken(gc *gin.Context) {
	cors.AllowCors(gc)
	var paramJwt request.RefreshTokenRequest
	if err := gc.ShouldBindHeader(&paramJwt); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.jwtService.RefreshToken(gc, paramJwt, repo.Timeout{})
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
		gc.Abort()
	}
	gc.Next()
}

// Authenticate godoc
// @Tags Token
// @Summary Provides a Sign In
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param company header string true "Company"
// @Param requestbody body request.TokenAWSRequest true "Request Body"
// @Success 200 {object} response.TokenAWSResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/signin [post]
func (c *jwtController) SigninAws(gc *gin.Context) {
	cors.AllowCors(gc)
	var bodyRequest request.TokenAWSRequest
	var headerRequest request.HeaderTokenRequest

	if err := gc.ShouldBindJSON(&bodyRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.ShouldBindHeader(&headerRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.jwtService.SigninAws(gc, bodyRequest, headerRequest, request.AwsRequest{})
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	gc.JSON(http.StatusOK, res)
}

// Authenticate godoc
// @Tags Token
// @Summary Provides a Refresh Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param company header string true "Company"
// @Param requestbody body request.RefreshTokenAWSRequest true "Request Body"
// @Success 200 {object} response.TokenAWSResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/resignin [post]
func (c *jwtController) ReSigninAws(gc *gin.Context) {
	cors.AllowCors(gc)
	var bodyRequest request.RefreshTokenAWSRequest
	var headerRequest request.HeaderTokenRequest

	if err := gc.ShouldBindJSON(&bodyRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.ShouldBindHeader(&headerRequest); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.jwtService.ReSigninAws(gc, bodyRequest, headerRequest, request.AwsRequest{})
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	gc.JSON(http.StatusOK, res)
}
