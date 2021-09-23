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
	DsfBranchCode       string
	Brand               string
	Model               string
	Variant             string
	DsfAssetCode        string
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
