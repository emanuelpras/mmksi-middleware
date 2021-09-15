package request

import (
	"middleware-mmksi/dsf/program/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderAssetCodeRequest struct {
	ApplicationName string `json:"applicationName"`
}

type AssetCodeRequest struct {
	VariantName      string
	CarCondition     string
	ManufacturedYear string
	ModelName        string
	BrandName        string
}

func (f *HeaderAssetCodeRequest) Validate() error {
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

func (f *AssetCodeRequest) Validate() error {
	if err := validation.Validate(f.VariantName, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "variant name not found",
				"id": "nama varian tidak ditemukan",
			},
		}
	}
	return nil
}
