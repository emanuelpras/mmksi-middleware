package service

import (
	"fmt"
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/service/request"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtService interface {
	CreateToken(gc *gin.Context, paramJwt request.TokenMmksiRequest)
	RefreshToken(gc *gin.Context, paramJwt request.TokenRefreshRequest)
}

type jwtService struct {
	jwtRepo repo.JwtRepo
}

func NewJwtService(
	jwtRepo repo.JwtRepo,
) JwtService {
	return &jwtService{
		jwtRepo: jwtRepo,
	}
}

func (s *jwtService) CreateToken(gc *gin.Context, paramJwt request.TokenMmksiRequest) {
	if paramJwt.Company != "" {
		if (paramJwt.Company == "mmksi") || (paramJwt.Company == "dsf") {
			result, err := s.jwtRepo.CreateToken(request.TokenMmksiRequest{
				Company: paramJwt.Company,
			})
			if err != nil {
				gc.JSON(http.StatusBadRequest, err)
			}
			gc.JSON(http.StatusOK, result)

		} else {
			gc.JSON(http.StatusUnauthorized, "unregistered company")
		}
	} else {
		if err := paramJwt.Validate(); err != nil {
			gc.JSON(http.StatusBadRequest, err)
		}
	}
}

func (s *jwtService) RefreshToken(gc *gin.Context, paramJwt request.TokenRefreshRequest) {
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
				res, err := s.jwtRepo.RefreshToken(request.TokenRefreshRequest{
					RefreshToken: str,
				})
				if err != nil {
					gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				gc.JSON(http.StatusOK, res)
				return
			}
			gc.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		gc.JSON(http.StatusBadRequest, "invalid token")
	} else {
		if err := paramJwt.Validate(); err != nil {
			gc.JSON(http.StatusBadRequest, err)
		}
	}
}
