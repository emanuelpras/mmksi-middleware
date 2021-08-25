package controller

import (
	"fmt"
	"net/http"
	"os"

	"middleware-mmksi/jwt/service"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
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
		gc.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	c.jwtService.CreateToken(gc, paramJwt)
}

func (c *jwtController) RefreshToken(gc *gin.Context) {
	var paramJwt request.TokenRefreshRequest
	if err := gc.ShouldBindHeader(&paramJwt); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	c.jwtService.RefreshToken(gc, paramJwt)
}

func (c *jwtController) Auth(gc *gin.Context) {
	option := os.Getenv("MIDDLEWARE_AUTH")
	if option == "YES" {
		tokenStringHeader := gc.Request.Header.Get("Auth")
		token, err := jwt.Parse(tokenStringHeader, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Method tidk diketahui atau bukan HS256 , method %V", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			gc.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(), "message": "Unauthorized",
			})
			gc.Abort()
		} else if token != nil {
			claims, _ := token.Claims.(jwt.MapClaims)
			if (claims["company"] == "dsf") || (claims["company"] == "mmksi") {
				gc.Next()
			} else {
				gc.JSON(http.StatusUnauthorized, gin.H{
					"error": "token invalid",
				})
				gc.Abort()
			}
		}
	} else {
		gc.Next()
	}
}
