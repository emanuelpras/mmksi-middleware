package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"middleware-mmksi/jwt/service"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtController struct {
	jwtService service.JwtService
}

type JwtController interface {
	GetFirstToken(context *gin.Context)
	Auth(context *gin.Context)
}

func NewJwtController(
	jwtService service.JwtService,
) *jwtController {
	return &jwtController{
		jwtService: jwtService,
	}
}

func (c *jwtController) GetFirstToken(gc *gin.Context) {

	var paramJwt request.FirstTokenRequest
	gc.BindHeader(&paramJwt)

	if (paramJwt.Company == "mmksi") || (paramJwt.Company == "dsf") {
		type authCustomClaims struct {
			Company string `json:"company"`
			jwt.StandardClaims
		}

		claims := &authCustomClaims{
			paramJwt.Company,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}

		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, _ := sign.SignedString([]byte("secret"))
		gc.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return

	} else {
		gc.JSON(http.StatusBadRequest, gin.H{
			"message": "company tidak terdaftar",
		})
		return
	}

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
