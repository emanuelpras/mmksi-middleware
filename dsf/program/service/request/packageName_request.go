package request

import (
	"middleware-mmksi/dsf/program/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderPackageNameRequest struct {
	ApplicationName string `json:"applicationName"`
	AssetCode       string `json:"assetCode"`
	BranchCode      string `json:"branchCode"`
}

type ParamsPackageNameRequest struct {
	CarCondition string `form:"carCondition"`
}

func (f *HeaderPackageNameRequest) Validate() error {
	if err := validation.Validate(f.ApplicationName, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "applicationName cannot be empty",
				"id": "applicationName harus diisi",
			},
		}
	}
	if err := validation.Validate(f.BranchCode, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "BranchCode cannot be empty",
				"id": "BranchCode harus diisi",
			},
		}
	}
	if err := validation.Validate(f.AssetCode, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "AssetCode cannot be empty",
				"id": "AssetCode harus diisi",
			},
		}
	}
	return nil
}
