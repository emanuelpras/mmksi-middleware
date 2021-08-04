package jwt

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/refactory-id/middleware-poc/response"
)

type FirstTokenRequest struct {
	Company string `form:"company"`
}

func (f *FirstTokenRequest) Validate() error {
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
