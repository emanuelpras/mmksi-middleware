package service

import (
	"math"
	"middleware-mmksi/soa/repo"
	"middleware-mmksi/soa/response"
	"middleware-mmksi/soa/service/request"
	"net/http"
)

type SoaService interface {
	VehicleMasterList(request request.SoaVehicleMasterRequest) (*response.ListVehicleMasterResponse, error)
}

type soaService struct {
	soaRepo repo.SoaRepo
}

func NewSoaService(soaRepo repo.SoaRepo) SoaService {
	return &soaService{
		soaRepo: soaRepo,
	}
}

func (s *soaService) VehicleMasterList(params request.SoaVehicleMasterRequest) (*response.ListVehicleMasterResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, counter, err := s.soaRepo.VehicleMasterList(repo.Pagination{
		Page:    params.Page,
		Limit:   params.Limit,
		Counter: ((params.Page - 1) * params.Limit),
	})

	if err != nil {
		response := &response.ErrorResponse{
			ErrorID: http.StatusUnprocessableEntity,
			Msg: map[string]string{
				"en": err.Error(),
				"id": "Entitas tidak dapat diproses",
			},
		}
		return nil, response
	}

	totalPage := math.Round(float64(counter) / float64(params.Limit))

	meta := response.Meta{
		Page:      int16(params.Page),
		Limit:     int16(params.Limit),
		TotalData: counter,
		TotalPage: int(totalPage),
	}

	res := response.ListVehicleMasterResponse{
		Meta: meta,
		Data: *result,
	}

	return &res, nil
}
