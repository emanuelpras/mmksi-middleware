package mmksi

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/refactory-id/middleware-poc/response"
)

type TokenRequest struct {
	Clientid   string `form:"clientid"`
	Dealercode string `form:"dealercode"`
	Username   string `form:"username"`
	Password   string `form:"password"`
}

func (f *TokenRequest) Validate() error {
	if err := validation.Validate(f.Clientid, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Client ID cannot be empty",
				"id": "Client ID harus diisi",
			},
		}
	}
	if err := validation.Validate(f.Dealercode, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Dealer Code cannot be empty",
				"id": "Dealer Code harus diisi",
			},
		}
	}
	if err := validation.Validate(f.Username, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Username cannot be empty",
				"id": "Username harus diisi",
			},
		}
	}
	if err := validation.Validate(f.Password, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Password cannot be empty",
				"id": "Password harus diisi",
			},
		}
	}
	return nil
}
