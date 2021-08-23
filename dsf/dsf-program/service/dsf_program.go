package service

import (
	"middleware-mmksi/dsf/dsf-program/repo"
	"middleware-mmksi/dsf/dsf-program/response"
)

type DsfProgramService interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
}

type dsfProgramService struct {
	dsfProgramRepo repo.DsfProgramRepo
}

func NewDsfProgramService(
	dsfProgramRepo repo.DsfProgramRepo,
) DsfProgramService {
	return &dsfProgramService{
		dsfProgramRepo: dsfProgramRepo,
	}
}

func (s *dsfProgramService) GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error) {

	result, err := s.dsfProgramRepo.GetAdditionalInsurance()
	if err != nil {
		return nil, err
	}

	return result, nil
}
