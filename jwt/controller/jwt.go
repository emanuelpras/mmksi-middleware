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
	if paramJwt.Company != "" {
		GenerateToken(gc)
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
	if paramJwt.Token != "" {
		GenerateRefreshToken(gc)
	} else {
		_, err := c.jwtService.RefreshToken(paramJwt)
		if err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func GenerateToken(gc *gin.Context) {
	var paramJwt request.TokenMmksiRequest
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

		claimsRefresh := &authCustomClaims{
			paramJwt.Company,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}

		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, _ := sign.SignedString([]byte("secret"))
		refresh := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claimsRefresh)
		tokenRefresh, _ := refresh.SignedString([]byte("secret"))
		gc.JSON(http.StatusOK, gin.H{
			"token":         token,
			"token refresh": tokenRefresh,
		})
	} else {
		gc.JSON(http.StatusBadRequest, gin.H{
			"message": "unregistered company",
		})
	}

}

func GenerateRefreshToken(gc *gin.Context) {
	tokenStringHeader := gc.Request.Header.Get("token")
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

		gc.BindHeader(&token)

		a := token.Claims.(jwt.MapClaims)
		b := a["company"]
		str := fmt.Sprintf("%v", b)
		type authCustomClaims struct {
			Company string `json:"company"`
			jwt.StandardClaims
		}

		if (str == "mmksi") || (str == "dsf") {

			claimsRefresh := &authCustomClaims{
				str,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
					IssuedAt:  time.Now().Unix(),
				},
			}
			refresh := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claimsRefresh)
			tokenRefresh, _ := refresh.SignedString([]byte("secret"))
			gc.JSON(http.StatusOK, gin.H{
				"token refresh": tokenRefresh,
			})

		} else {
			gc.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalid",
			})
			gc.Abort()
		}
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
