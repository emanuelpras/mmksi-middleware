package service

import (
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"
)

type JwtService interface {
	CreateToken(params request.FirstTokenRequest) (*response.FirtsTokenResponse, error)
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

func (s *jwtService) CreateToken(params request.FirstTokenRequest) (*response.FirtsTokenResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	result, err := s.jwtRepo.CreateToken(repo.ParamToken{
		Company: params.Company,
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
