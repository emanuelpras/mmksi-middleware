package request

import (
	"middleware-mmksi/jwt/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TokenAWSRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenMmksiRequest struct {
	Company string `form:"company"`
}

type TokenRefreshRequest struct {
	RefreshToken string `form:"refreshToken"`
}

type AwsRequest struct {
	Region       string
	ClientID     string
	ClientSecret string
}

func (f *TokenAWSRequest) Validate() error {

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

func (f *TokenMmksiRequest) Validate() error {
	if err := validation.Validate(f.Company, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Company cannot be empty",
				"id": "Company harus diisi",
			},
		}
	}
	return nil
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
