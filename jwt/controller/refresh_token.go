package controller

import (
	"fmt"
	"net/http"

	"middleware-mmksi/jwt/service"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type refreshTokenController struct {
	refreshTokenService service.RefreshTokenService
}

type RefreshTokenController interface {
	RefreshToken(context *gin.Context)
}

func NewRefreshTokenController(
	refreshTokenService service.RefreshTokenService,
) *refreshTokenController {
	return &refreshTokenController{
		refreshTokenService: refreshTokenService,
	}
}

func (c *refreshTokenController) RefreshToken(gc *gin.Context) {
	var paramJwt request.TokenRefreshRequest
	if err := gc.ShouldBindHeader(&paramJwt); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	res, err := c.refreshTokenService.ValidateRefreshToken(paramJwt)

	if paramJwt.RefreshToken != "" {

		token, _ := jwt.Parse(paramJwt.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if (claims["company"] == "mmksi") || (claims["company"] == "dsf") {
				company := claims["company"]
				str := fmt.Sprintf("%v", company)
				res, err := c.refreshTokenService.RefreshToken(str)
				if err != nil {
					gc.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
					return
				}

				gc.JSON(http.StatusOK, res)
				return
			}
			gc.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		gc.JSON(http.StatusBadRequest, "invalid token")
	} else if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		gc.JSON(http.StatusOK, res)
	}

}
