package controller

import (
	"fmt"
	"log"
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
	CreateToken(context *gin.Context)
	RefreshToken(context *gin.Context)
	Auth(context *gin.Context)
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
	gc.BindHeader(&paramJwt)
	log.Print("com", paramJwt)

	if paramJwt.Company != "" {
		if (paramJwt.Company == "mmksi") || (paramJwt.Company == "dsf") {
			company := paramJwt.Company
			tokens, err := GenerateToken(gc, company)
			if err != nil {
				gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			gc.JSON(http.StatusOK, tokens)
			return
		}
		gc.JSON(http.StatusUnauthorized, "unregistered company")
	} else {
		_, err := c.jwtService.CreateToken(paramJwt)
		if err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func (c *jwtController) RefreshToken(gc *gin.Context) {
	var paramJwt request.TokenRefreshRequest
	gc.BindHeader(&paramJwt)
	log.Print("tok", paramJwt)

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
				log.Print(str)
				newToken, err := GenerateToken(gc, str)
				if err != nil {
					gc.JSON(http.StatusBadRequest, err)
					return
				}
				gc.JSON(http.StatusOK, newToken)
				return
			}
			gc.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		gc.JSON(http.StatusBadRequest, "invalid token")
	} else {
		_, err := c.jwtService.RefreshToken(paramJwt)
		if err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func GenerateToken(gc *gin.Context, company string) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["company"] = company
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["company"] = company
	rtClaims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	refreshToken, err := refresh.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
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
