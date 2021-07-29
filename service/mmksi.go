package service

import (
	"github.com/refactory-id/middleware-poc/repo"
	response "github.com/refactory-id/middleware-poc/response/mmksi"
	"github.com/refactory-id/middleware-poc/service/mmksi"
)

type MmksiService interface {
	GetToken(params mmksi.TokenRequest) (*response.TokenResponse, error)
	GetVehicles(params mmksi.VehicleRequest, authorizationMmksi mmksi.VehicleRequestAuthorization) (*response.VehicleResponse, error)
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

func (s *mmksiService) GetVehicles(params mmksi.VehicleRequest, authorizationMmksi mmksi.VehicleRequestAuthorization) (*response.VehicleResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mmksiRepo.GetVehicles(repo.GetVehicleParams{
		Page: params.Page,
	}, repo.GetVehicleHeaderAuthorization{Authorization: authorizationMmksi.Authorization})

	if err != nil {
		return nil, err
	}

	return result, nil
}
