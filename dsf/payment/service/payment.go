package service

import (
	"middleware-mmksi/dsf/payment/repo"
	"middleware-mmksi/dsf/payment/response"
)

type DsfProgramService interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
	GetPackageNames() (*response.GetPackageNames, error)
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

func (s *dsfProgramService) GetPackageNames() (*response.GetPackageNames, error) {

	result, err := s.dsfProgramRepo.GetPackageNames()
	if err != nil {
		return nil, err
	}

	return result, nil
}
