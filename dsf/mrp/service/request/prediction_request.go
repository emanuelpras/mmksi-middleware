package request

import (
	er "middleware-mmksi/dsf/mrp/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PredictionRequest struct {
	Brand        string `json:"BRAND"`
	Model        string `json:"MODEL"`
	Variant      string `json:"VARIANT"`
	Year         int16  `json:"YEAR"`
	Distance     int64  `json:"DISTANCE"`
	Transmission string `json:"TRANSMISSION"`
	Color        string `json:"COLOR"`
	SellerType   string `json:"TIPE_PENJUAL"`
	City         string `json:"CITY"`
	Province     string `json:"PROVINCE"`
	Company      string `json:"COMPANY"`
}

func (f *PredictionRequest) Validate() error {
	if err := validation.Validate(f.Brand, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Brand cannot be empty",
				"id": "Brand harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Model, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Model cannot be empty",
				"id": "Model harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Variant, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Variant cannot be empty",
				"id": "Variant harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Year, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Year cannot be empty",
				"id": "Year harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Distance, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Distance cannot be empty",
				"id": "Distance harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Transmission, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Transmission cannot be empty",
				"id": "Transmission harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Color, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Color cannot be empty",
				"id": "Color harus diisi",
			},
		}
	}

	if err := validation.Validate(f.SellerType, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Seller Type cannot be empty",
				"id": "Tipe Penjual harus diisi",
			},
		}
	}

	if err := validation.Validate(f.SellerType, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Seller Type cannot be empty",
				"id": "Tipe Penjual harus diisi",
			},
		}
	}

	if err := validation.Validate(f.City, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "City cannot be empty",
				"id": "City harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Province, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Province cannot be empty",
				"id": "Province harus diisi",
			},
		}
	}

	if err := validation.Validate(f.Company, validation.Required); err != nil {
		return &er.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Company cannot be empty",
				"id": "Company harus diisi",
			},
		}
	}

	return nil
}
