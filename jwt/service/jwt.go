package service

import (
	"fmt"
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtService interface {
	CreateToken(gc *gin.Context, paramJwt request.TokenMmksiRequest) (*response.TokenMmksiResponse, error)
	RefreshToken(gc *gin.Context, paramJwt request.TokenRefreshRequest) (*response.TokenMmksiResponse, error)
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

func (s *jwtService) CreateToken(gc *gin.Context, paramJwt request.TokenMmksiRequest) (*response.TokenMmksiResponse, error) {
	if err := paramJwt.Validate(); err != nil {
		return nil, err
	}
	result, err := s.jwtRepo.CreateToken(request.TokenMmksiRequest{
		Company: paramJwt.Company,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *jwtService) RefreshToken(gc *gin.Context, paramJwt request.TokenRefreshRequest) (*response.TokenMmksiResponse, error) {
	if err := paramJwt.Validate(); err != nil {
		return nil, err
	}
	token, _ := jwt.Parse(paramJwt.RefreshToken, func(token *jwt.Token) (interface{}, error) {
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
				return nil, err
			}
			return res, err
		}
		return nil, &response.ErrorResponse{
			ErrorID: 400,
			Msg: map[string]string{
				"en": "Company unregistered",
				"id": "Company tidak terdaftar",
			},
		}
	}
	return nil, &response.ErrorResponse{
		ErrorID: 400,
		Msg: map[string]string{
			"en": "Invalid token or token expired",
			"id": "Token tidak valid atau token telah kadaluarsa",
		},
	}
}
