package request

import (
	"middleware-mmksi/jwt/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AuthRequest struct {
	Auth string `form:"Auth"`
}

func (f *AuthRequest) Validate() error {
	if err := validation.Validate(f.Auth, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 400,
			Msg: map[string]string{
				"en": "Auth cannot be empty",
				"id": "Auth harus diisi",
			},
		}
	}
	return nil
}
