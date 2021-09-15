package request

import (
	"middleware-mmksi/dsf/mrp/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderTenorRequest struct {
	ApplicationName string `json:"applicationName"`
}

type TenorRequest struct {
	UnitPrice           interface{}
	Province            interface{}
	City                interface{}
	DsfBranchCode       interface{}
	Brand               interface{}
	Model               interface{}
	Variant             interface{}
	DsfAssetCode        interface{}
	CarCondition        interface{}
	ManufacturedYear    interface{}
	LoanPackageName     interface{}
	Caroserie           interface{}
	PaymentType         interface{}
	Insurances          insurances
	Fee                 fee
	ProvisionPercentage interface{}
	TenorInMonths       interface{}
	SimulationType      interface{}
	SimulationValue     interface{}
}

type insurances struct {
	InsuranceType            interface{}
	AdditionalInsurances     interface{}
	LifeInsurance            interface{}
	TanggungJawabPihakKetiga tanggungJawabPihakKetiga
	PutAsOnLoan              interface{}
}

type tanggungJawabPihakKetiga struct {
	IsApplied         interface{}
	UangPertanggungan interface{}
}

type fee struct {
	BeaPolis interface{}
	AdminFee interface{}
}

func (f *HeaderTenorRequest) Validate() error {
	if err := validation.Validate(f.ApplicationName, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "applicationName cannot be empty",
				"id": "applicationName harus diisi",
			},
		}
	}
	return nil
}
