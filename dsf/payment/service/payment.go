package service

import (
	"middleware-mmksi/dsf/payment/repo"
	"middleware-mmksi/dsf/payment/response"
	"middleware-mmksi/dsf/payment/service/request"
)

type DsfPaymentService interface {
	GetTenor(params request.HeaderTenorRequest, reqBody request.TenorRequest) (*response.TenorResponse, error)
}

type dsfPaymentService struct {
	dsfPaymentRepo repo.DsfPaymentRepo
}

func NewDsfPaymentService(
	dsfPaymentRepo repo.DsfPaymentRepo,
) DsfPaymentService {
	return &dsfPaymentService{
		dsfPaymentRepo: dsfPaymentRepo,
	}
}

func (s *dsfPaymentService) GetTenor(params request.HeaderTenorRequest, reqBody request.TenorRequest) (*response.TenorResponse, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfPaymentRepo.GetTenor(request.HeaderTenorRequest{
		ApplicationName: params.ApplicationName,
	}, request.TenorRequest{
		UnitPrice:           reqBody.UnitPrice,
		Province:            reqBody.Province,
		City:                reqBody.City,
		DsfBranchCode:       reqBody.DsfBranchCode,
		Brand:               reqBody.Brand,
		Model:               reqBody.Model,
		Variant:             reqBody.Variant,
		DsfAssetCode:        reqBody.DsfAssetCode,
		CarCondition:        reqBody.CarCondition,
		ManufacturedYear:    reqBody.ManufacturedYear,
		LoanPackageName:     reqBody.LoanPackageName,
		Caroserie:           reqBody.Caroserie,
		PaymentType:         reqBody.PaymentType,
		Insurances:          reqBody.Insurances,
		Fee:                 reqBody.Fee,
		ProvisionPercentage: reqBody.ProvisionPercentage,
		TenorInMonths:       reqBody.TenorInMonths,
		SimulationType:      reqBody.SimulationType,
		SimulationValue:     reqBody.SimulationValue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
