package request

import (
	"middleware-mmksi/dsf/payment/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderPackageRequest struct {
	ApplicationName string `json:"applicationName"`
}

type PackageRequest struct {
	Brand        string
	Model        string
	Variant      string
	Province     string
	City         string
	PackageName  string
	CarCondition string
}

func (f *HeaderPackageRequest) Validate() error {
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

func (f *PackageRequest) Validate() error {
	if err := validation.Validate(f.Variant, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "variant not found",
				"id": "varian tidak ditemukan",
			},
		}
	}

	if err := validation.Validate(f.CarCondition, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "car condition not found",
				"id": "kondisi mobil tidak ditemukan",
			},
		}
	}
	return nil
}
