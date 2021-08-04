package mmksi

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/refactory-id/middleware-poc/response"
)

type VehicleRequest struct {
	Page int64 `json:"pages"`
}

type VehicleRequestAuthorization struct {
	AccessToken string `form:"AccessToken"`
	TokenType   string `form:"TokenType"`
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
