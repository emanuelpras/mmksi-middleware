package service

import (
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"
	"os"

	"github.com/gin-gonic/gin"
)

type JwtService interface {
	CreateToken(gc *gin.Context, paramJwt request.TokenMmksiRequest, timeout repo.Timeout) (*response.TokenMmksiResponse, error)
	RefreshToken(gc *gin.Context, paramJwt request.TokenRefreshRequest, timeout repo.Timeout) (*response.TokenMmksiResponse, error)
	Auth(gc *gin.Context, auth request.AuthRequest) error
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

func (s *jwtService) CreateToken(gc *gin.Context, paramJwt request.TokenMmksiRequest, timeout repo.Timeout) (*response.TokenMmksiResponse, error) {
	if err := paramJwt.Validate(); err != nil {
		return nil, err
	}

	result, err := s.jwtRepo.CreateToken(request.TokenMmksiRequest{
		Company: paramJwt.Company,
	}, repo.Timeout{
		TimeoutAccessToken:  os.Getenv("ACCESS_TOKEN_TIMEOUT"),
		TimeoutRefreshToken: os.Getenv("REFRESH_TOKEN_TIMEOUT"),
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *jwtService) RefreshToken(gc *gin.Context, paramJwt request.TokenRefreshRequest, timeout repo.Timeout) (*response.TokenMmksiResponse, error) {
	if err := paramJwt.Validate(); err != nil {
		return nil, err
	}
	result, err := s.jwtRepo.RefreshToken(request.TokenRefreshRequest{
		RefreshToken: paramJwt.RefreshToken,
	}, repo.Timeout{
		TimeoutAccessToken:  os.Getenv("ACCESS_TOKEN_TIMEOUT"),
		TimeoutRefreshToken: os.Getenv("REFRESH_TOKEN_TIMEOUT"),
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *jwtService) Auth(gc *gin.Context, auth request.AuthRequest) error {
	option := os.Getenv("MIDDLEWARE_AUTH")

	if option == "YES" {
		if err := auth.Validate(); err != nil {
			return err
		}
		err := s.jwtRepo.Auth(gc, request.AuthRequest{
			Auth: auth.Auth,
		}, repo.Timeout{})
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
