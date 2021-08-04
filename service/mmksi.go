package service

import (
	"os"

	"github.com/refactory-id/middleware-poc/repo"
	response "github.com/refactory-id/middleware-poc/response/mmksi"
	"github.com/refactory-id/middleware-poc/service/mmksi"
)

type MmksiService interface {
	GetToken(params mmksi.TokenRequest) (*response.TokenResponse, error)
	GetVehicle(params mmksi.VehicleRequest, authorizationMmksi mmksi.VehicleRequestAuthorization) (*response.VehicleResponse, error)
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

	result, err := s.mmksiRepo.GetToken(repo.GetTokenParams{
		Clientid:   os.Getenv("CLIENT_ID"),
		Dealercode: os.Getenv("DEALER_CODE"),
		Username:   os.Getenv("USERNAME"),
		Password:   os.Getenv("PASSWORD"),
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mmksiService) GetVehicle(params mmksi.VehicleRequest, authorizationMmksi mmksi.VehicleRequestAuthorization) (*response.VehicleResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mmksiRepo.GetVehicle(repo.GetVehicleParams{
		Page: params.Page,
	}, repo.GetHeaderAuthorization{
		AccessToken: authorizationMmksi.AccessToken,
		TokenType:   authorizationMmksi.TokenType,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
