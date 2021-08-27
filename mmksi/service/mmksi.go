package service

import (
	"os"

	"middleware-mmksi/mmksi/repo"
	"middleware-mmksi/mmksi/response"
	"middleware-mmksi/mmksi/service/request"
)

type MmksiService interface {
	GetToken(params request.TokenRequest) (*response.TokenResponse, error)
	GetVehicle(params request.VehicleRequest, authorizationMmksi request.VehicleRequestAuthorization) (*response.VehicleResponse, error)
	GetVehicleColor(params request.VehicleRequest, authorizationMmksi request.VehicleRequestAuthorization) (*response.VehicleColorResponse, error)
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

func (s *mmksiService) GetToken(params request.TokenRequest) (*response.TokenResponse, error) {

	result, err := s.mmksiRepo.GetToken(request.TokenRequest{
		Clientid:   os.Getenv("DNET_CLIENT_ID"),
		Dealercode: os.Getenv("DNET_DEALER_CODE"),
		Username:   os.Getenv("DNET_USERNAME"),
		Password:   os.Getenv("DNET_PASSWORD"),
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mmksiService) GetVehicle(params request.VehicleRequest, authorizationMmksi request.VehicleRequestAuthorization) (*response.VehicleResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mmksiRepo.GetVehicle(request.VehicleRequest{
		Page:        params.Page,
		FindVehicle: params.FindVehicle,
		SortVehicle: params.SortVehicle,
	}, request.VehicleRequestAuthorization{
		AccessToken: authorizationMmksi.AccessToken,
		TokenType:   authorizationMmksi.TokenType,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mmksiService) GetVehicleColor(params request.VehicleRequest, authorizationMmksi request.VehicleRequestAuthorization) (*response.VehicleColorResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mmksiRepo.GetVehicleColor(request.VehicleRequest{
		Page:        params.Page,
		FindVehicle: params.FindVehicle,
		SortVehicle: params.SortVehicle,
	}, request.VehicleRequestAuthorization{
		AccessToken: authorizationMmksi.AccessToken,
		TokenType:   authorizationMmksi.TokenType,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
