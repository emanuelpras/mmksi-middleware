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
	RefreshToken(gc *gin.Context, paramJwt request.RefreshTokenRequest, timeout repo.Timeout) (*response.TokenMmksiResponse, error)
	Auth(gc *gin.Context, auth request.AuthRequest) error
	SigninAws(gc *gin.Context, param request.TokenAWSRequest, config request.AwsRequest) (*response.TokenAWSResponse, error)
	ReSigninAws(gc *gin.Context, param request.RefreshTokenAWSRequest, config request.AwsRequest) (*response.RefreshTokenAWSResponse, error)
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

func (s *jwtService) RefreshToken(gc *gin.Context, paramJwt request.RefreshTokenRequest, timeout repo.Timeout) (*response.TokenMmksiResponse, error) {
	if err := paramJwt.Validate(); err != nil {
		return nil, err
	}
	result, err := s.jwtRepo.RefreshToken(request.RefreshTokenRequest{
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

func (s *jwtService) SigninAws(gc *gin.Context, param request.TokenAWSRequest, config request.AwsRequest) (*response.TokenAWSResponse, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	result, err := s.jwtRepo.SigninAws(request.TokenAWSRequest{
		Username: param.Username,
		Password: param.Password,
	}, request.AwsRequest{
		Region:       os.Getenv("REGION_AWS"),
		ClientID:     os.Getenv("CLIENT_ID_AWS"),
		ClientSecret: os.Getenv("CLIENT_SECRET_AWS"),
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *jwtService) ReSigninAws(gc *gin.Context, param request.RefreshTokenAWSRequest, config request.AwsRequest) (*response.RefreshTokenAWSResponse, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	result, err := s.jwtRepo.ReSigninAws(request.RefreshTokenAWSRequest{
		Username:     param.Username,
		RefreshToken: param.RefreshToken,
	}, request.AwsRequest{
		Region:       os.Getenv("REGION_AWS"),
		ClientID:     os.Getenv("CLIENT_ID_AWS"),
		ClientSecret: os.Getenv("CLIENT_SECRET_AWS"),
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
