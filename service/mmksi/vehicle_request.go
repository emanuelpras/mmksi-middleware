package mmksi

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/refactory-id/middleware-poc/response"
)

type VehicleRequest struct {
	Page int64 `json:"pages"`
}

type VehicleRequestAuthorization struct {
	Authorization string `form:"Authorization"`
}

func (f *VehicleRequest) Validate() error {
	if err := validation.Validate(f.Page, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "page not found",
				"id": "halaman tidak ditemukan",
			},
		}
	}
	return nil
}

func (f *VehicleRequestAuthorization) Validate() error {
	if err := validation.Validate(f.Authorization, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 401,
			Msg: map[string]string{
				"en": "authorization not allowed",
				"id": "otorisasi tidak diizinkan",
			},
		}
	}
	return nil
}
