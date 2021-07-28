package service

import (
	"github.com/refactory-id/middleware-poc/repo"
	"github.com/refactory-id/middleware-poc/response"
	"github.com/refactory-id/middleware-poc/service/mmksi"
)

type MmksiService interface {
	GetToken(params mmksi.TokenRequest) (*response.TokenResponse, error)
}

type mmksiService struct {
	mmksiRepo repo.MmksiRepo
}

func NewMmksiService(
	mmksiRepo repo.MmksiRepo,
) MmksiService {
	return &mmksiService{
		mmksiRepo: mmksiRepo,
	}
}

func (s *mmksiService) GetToken(params mmksi.TokenRequest) (*response.TokenResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mmksiRepo.GetToken(repo.GetTokenParams{
		Clientid:   params.Clientid,
		Dealercode: params.Dealercode,
		Username:   params.Username,
		Password:   params.Password,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
