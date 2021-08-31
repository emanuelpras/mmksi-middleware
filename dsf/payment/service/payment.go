package service

import (
	"middleware-mmksi/dsf/payment/repo"
	"middleware-mmksi/dsf/payment/response"
	"middleware-mmksi/dsf/payment/service/request"
)

type DsfProgramService interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
	GetPackageNames() (*response.PackageNameResponse, error)
	GetCarConditions() (*response.CarConditionResponse, error)
	GetPackages(params request.HeaderPackageRequest) (*response.PackageResponse, error)
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

func (s *dsfProgramService) GetPackageNames() (*response.PackageNameResponse, error) {

	result, err := s.dsfProgramRepo.GetPackageNames()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetCarConditions() (*response.CarConditionResponse, error) {

	result, err := s.dsfProgramRepo.GetCarConditions()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetPackages(params request.HeaderPackageRequest) (*response.PackageResponse, error) {
	// if err := params.Validate(); err != nil {
	// 	return nil, err
	// }

	result, err := s.dsfProgramRepo.GetPackages(request.HeaderPackageRequest{
		ApplicationName: params.ApplicationName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
