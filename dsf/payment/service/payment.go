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
	GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error)
	GetUnitByModels(paramHeader request.HeaderUnitByModelsRequest) (*response.UnitByModelsResponse, error)
	GetPaymentTypes() (*response.PaymentTypesResponse, error)
	GetBranchID() (*response.BranchResponse, error)
	GetInsuranceTypes() (*response.InsuranceTypesResponse, error)
	GetInsurance(params request.InsuranceRequest) (*response.InsuranceResponse, error)
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

func (s *dsfProgramService) GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error) {
	if err := paramHeader.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetPackages(request.HeaderPackageRequest{
		ApplicationName: paramHeader.ApplicationName,
	}, request.PackageRequest{
		Brand:        reqBody.Brand,
		Model:        reqBody.Model,
		Variant:      reqBody.Variant,
		Province:     reqBody.Province,
		City:         reqBody.City,
		PackageName:  reqBody.PackageName,
		CarCondition: reqBody.CarCondition,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetUnitByModels(paramHeader request.HeaderUnitByModelsRequest) (*response.UnitByModelsResponse, error) {
	if err := paramHeader.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetUnitByModels(request.HeaderUnitByModelsRequest{
		ApplicationName: paramHeader.ApplicationName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetPaymentTypes() (*response.PaymentTypesResponse, error) {

	result, err := s.dsfProgramRepo.GetPaymentTypes()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetBranchID() (*response.BranchResponse, error) {

	result, err := s.dsfProgramRepo.GetBranchID()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetInsuranceTypes() (*response.InsuranceTypesResponse, error) {

	result, err := s.dsfProgramRepo.GetInsuranceTypes()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetInsurance(params request.InsuranceRequest) (*response.InsuranceResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetInsurance(request.InsuranceRequest{
		DsfBranchId:       params.DsfBranchId,
		VehicleCategory:   params.VehicleCategory,
		InsuranceTypeCode: params.InsuranceTypeCode,
		CarCondition:      params.CarCondition,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
