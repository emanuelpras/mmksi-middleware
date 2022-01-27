package request

import (
	er "middleware-mmksi/soa/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type SoaVehicleMasterRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type VehicleMasterByAssetCodeRequest struct {
	AssetCode string `query:"assetCode"`
}

func (f *SoaVehicleMasterRequest) Validate() error {
	if err := validation.Validate(f.Page, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Page cannot be empty",
				"id": "Page harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Limit, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Limit cannot be empty",
				"id": "Limit harus diisi",
			},
		}
	}

	return nil
}
