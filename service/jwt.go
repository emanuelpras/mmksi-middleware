package service

import (
	"github.com/refactory-id/middleware-poc/repo"
	response "github.com/refactory-id/middleware-poc/response/jwt"

	"github.com/refactory-id/middleware-poc/service/jwt"
)

type JwtService interface {
	GetFirstToken(params jwt.FirstTokenRequest) (*response.FirtsTokenResponse, error)
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

func (s *jwtService) GetFirstToken(params jwt.FirstTokenRequest) (*response.FirtsTokenResponse, error) {
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
