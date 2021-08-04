package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/refactory-id/middleware-poc/service"
	serviceJwt "github.com/refactory-id/middleware-poc/service/jwt"
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

	var paramJwt serviceJwt.FirstTokenRequest
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

		if err := godotenv.Load(".env"); err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{
				"message": ".env ngga ketemu",
			})
			return
		} else {
			sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
			token, _ := sign.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
			gc.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	} else {
		gc.JSON(http.StatusBadRequest, gin.H{
			"message": "company ngga kedaftar",
		})
		return
	}

}

func (c *jwtController) Auth(gc *gin.Context) {
	secret := os.Getenv("SECRET_TOKEN")
	tokenStringHeader := gc.Request.Header.Get("MmksiAuth")
	token, err := jwt.Parse(tokenStringHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method tidk diketahui atau bukan HS256 , method %V", token.Header["alg"])
		}
		return []byte(secret), nil
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
}
