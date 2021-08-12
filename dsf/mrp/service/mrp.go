package service

import (
	"middleware-mmksi/dsf/mrp/repo"
	"middleware-mmksi/dsf/mrp/response"
	"middleware-mmksi/dsf/mrp/service/request"
)

type MrpService interface {
	GetVehicles(params request.GetVehicleRequest) (*response.GetVehiclesResponse, error)
	GetRegions(params request.GetRegionsRequest) (*response.GetRegionsResponse, error)
	GetPrediction(params request.PredictionRequest) (*response.PredictionResponse, error)
}

type mrpService struct {
	mrpRepo repo.MrpRepo
}

func NewMrpService(
	mrpRepo repo.MrpRepo,
) MrpService {
	return &mrpService{
		mrpRepo: mrpRepo,
	}
}

func (s *mrpService) GetVehicles(params request.GetVehicleRequest) (*response.GetVehiclesResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mrpRepo.GetVehicles(repo.GetVehiclesParams{
		BrandName: params.BrandName,
		ModelName: params.ModelName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mrpService) GetRegions(params request.GetRegionsRequest) (*response.GetRegionsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mrpRepo.GetRegions(repo.GetRegionsParams{
		Province: params.Province,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mrpService) GetPrediction(params request.PredictionRequest) (*response.PredictionResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.mrpRepo.GetPrediction(repo.GetPredictionParams{
		Brand:        params.Brand,
		Model:        params.Model,
		Variant:      params.Variant,
		Year:         params.Year,
		Distance:     params.Distance,
		Transmission: params.Transmission,
		Color:        params.Color,
		SellerType:   params.SellerType,
		City:         params.City,
		Province:     params.Province,
		Company:      params.Company,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
