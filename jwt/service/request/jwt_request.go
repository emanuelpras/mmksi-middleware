package request

import (
	"middleware-mmksi/jwt/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TokenMmksiRequest struct {
	Company string `form:"company"`
}

type TokenRefreshRequest struct {
	RefreshToken string `form:"refreshToken"`
}

func (f *TokenMmksiRequest) Validate() error {
	if err := validation.Validate(f.Company, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Company cannot be empty",
				"id": "Company harus diisi",
			},
		}
	} else if f.Company == "mmksi" || f.Company == "dsf" {
		return nil
	}
	return &response.ErrorResponse{
		ErrorID: 400,
		Msg: map[string]string{
			"en": "Company unregistered",
			"id": "Company tidak terdaftar",
		},
	}
}

func (f *TokenRefreshRequest) Validate() error {
	if err := validation.Validate(f.RefreshToken, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Refresh Token cannot be empty",
				"id": "Refresh Token harus diisi",
			},
		}
	}
	return nil
}
