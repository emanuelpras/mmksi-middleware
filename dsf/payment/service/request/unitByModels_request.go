package request

import (
	"middleware-mmksi/dsf/payment/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderUnitByModelsRequest struct {
	ApplicationName string `json:"applicationName"`
}

func (f *HeaderUnitByModelsRequest) Validate() error {
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
