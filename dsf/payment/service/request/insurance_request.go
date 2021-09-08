package request

import (
	"middleware-mmksi/dsf/payment/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type InsuranceRequest struct {
	DsfBranchId       string `form:"DsfBranchId"`
	VehicleCategory   string `form:"VehicleCategory"`
	InsuranceTypeCode string `form:"InsuranceTypeCode"`
	CarCondition      string `form:"CarCondition"`
}

func (f *InsuranceRequest) Validate() error {
	if err := validation.Validate(f.DsfBranchId, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "DsfBranchId cannot be empty",
				"id": "DsfBranchId harus diisi",
			},
		}
	}
	if err := validation.Validate(f.VehicleCategory, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "VehicleCategory cannot be empty",
				"id": "VehicleCategory harus diisi",
			},
		}
	}
	if err := validation.Validate(f.InsuranceTypeCode, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "InsuranceTypeCode cannot be empty",
				"id": "InsuranceTypeCode harus diisi",
			},
		}
	}
	if err := validation.Validate(f.CarCondition, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "CarCondition cannot be empty",
				"id": "CarCondition harus diisi",
			},
		}
	}
	return nil
}
