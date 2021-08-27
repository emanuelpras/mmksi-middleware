package request

import (
	"middleware-mmksi/mmksi/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type VehicleRequest struct {
	Page int64 `json:"pages"`

	FindVehicle []findVehicleRequest `json:"find"`
	SortVehicle []sortVehicleRequest `json:"sort"`
}

type findVehicleRequest struct {
	MatchType     int64  `json:"MatchType"`
	PropertyName  string `json:"PropertyName"`
	PropertyValue string `json:"PropertyValue"`
	SqlOperation  int64  `json:"SqlOperation"`
}

type sortVehicleRequest struct {
	SortColumn    string `json:"SortColumn"`
	SortDirection int64  `json:"SortDirection"`
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
