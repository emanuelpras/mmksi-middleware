package service

import (
	"middleware-mmksi/jwt/repo"
	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"
)

type JwtService interface {
	GetFirstToken(params request.FirstTokenRequest) (*response.FirtsTokenResponse, error)
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

func (s *jwtService) GetFirstToken(params request.FirstTokenRequest) (*response.FirtsTokenResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	result, err := s.jwtRepo.GetFirstToken(repo.ParamToken{
		Company: params.Company,
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
