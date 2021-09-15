package request

type CitiesRequest struct {
	Search       string `form:"search"`
	ProvinceCode string `form:"province_code"`
	Offset       int    `form:"offset"`
	Limit        int    `form:"limit"`
}

func (f *CitiesRequest) Validate() error {
	return nil
}
