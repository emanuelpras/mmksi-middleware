package service

import (
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RefreshTokenService interface {
	ValidateRefreshToken(paramJwt request.TokenRefreshRequest) (*response.TokenMmksiResponse, error)
	RefreshToken(company string) (*response.TokenMmksiResponse, error)
}

type refreshTokenService struct {
	refreshTokenRepo repo.RefreshTokenRepo
}

func NewRefreshTokenService(
	refreshTokenRepo repo.RefreshTokenRepo,
) RefreshTokenService {
	return &refreshTokenService{
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *refreshTokenService) RefreshToken(company string) (*response.TokenMmksiResponse, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["company"] = company
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return &response.TokenMmksiResponse{}, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["company"] = company
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	refreshToken, err2 := refresh.SignedString([]byte("secret"))
	if err2 != nil {
		return &response.TokenMmksiResponse{}, err2
	}

	result, err := s.refreshTokenRepo.RefreshToken(accessToken, refreshToken)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *refreshTokenService) ValidateRefreshToken(paramJwt request.TokenRefreshRequest) (*response.TokenMmksiResponse, error) {

	if err := paramJwt.Validate(); err != nil {
		return nil, err
	}

	result, err := s.refreshTokenRepo.ValidateRefreshToken(repo.ParamRefreshToken(paramJwt))

	if err != nil {
		return nil, err
	}
	return result, nil
}
