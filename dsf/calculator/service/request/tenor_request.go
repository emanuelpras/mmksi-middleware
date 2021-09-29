package request

import (
	"middleware-mmksi/dsf/calculator/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderTenorRequest struct {
	ApplicationName string `json:"applicationName"`
}

type TenorRequest struct {
	UnitPrice           int
	Province            string
	City                string
	Brand               string
	Model               string
	Variant             string
	CarCondition        string
	ManufacturedYear    string
	LoanPackageName     string
	Caroserie           int
	PaymentType         string
	Insurances          insurances
	Fee                 fee
	ProvisionPercentage int
	TenorInMonths       int
	SimulationType      string
	SimulationValue     int
}

type AllTenorRequest struct {
	UnitPrice           int
	Province            string
	City                string
	Brand               string
	Model               string
	Variant             string
	CarCondition        string
	ManufacturedYear    string
	LoanPackageName     string
	Caroserie           int
	PaymentType         string
	Insurances          insurances
	Fee                 fee
	ProvisionPercentage int
	SimulationType      string
	SimulationValue     int
}

type insurances struct {
	InsuranceType            string
	AdditionalInsurances     []string
	LifeInsurance            bool
	TanggungJawabPihakKetiga tanggungJawabPihakKetiga
	PutAsOnLoan              bool
}

type tanggungJawabPihakKetiga struct {
	IsApplied         bool
	UangPertanggungan int
}

type fee struct {
	BeaPolis int
	AdminFee int
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

func (f *TenorRequest) Validate() error {
	if err := validation.Validate(f.UnitPrice, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "unit price cannot be empty",
				"id": "unit price harus diisi",
			},
		}
	}
	if err := validation.Validate(f.City, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "city cannot be empty",
				"id": "city harus diisi",
			},
		}
	}
	if err := validation.Validate(f.Variant, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "variant cannot be empty",
				"id": "variant harus diisi",
			},
		}
	}
	if err := validation.Validate(f.TenorInMonths, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "tenor in months cannot be empty",
				"id": "tenor in months harus diisi",
			},
		}
	}
	if err := validation.Validate(f.SimulationValue, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "simulation value cannot be empty",
				"id": "simulation value harus diisi",
			},
		}
	}
	return nil
}

func (f *AllTenorRequest) Validate() error {
	if err := validation.Validate(f.UnitPrice, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "unit price cannot be empty",
				"id": "unit price harus diisi",
			},
		}
	}
	if err := validation.Validate(f.SimulationValue, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "simulation value cannot be empty",
				"id": "simulation value harus diisi",
			},
		}
	}
	return nil
}
