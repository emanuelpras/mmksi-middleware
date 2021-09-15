package request

type ProvincesRequest struct {
	Search string `form:"search"`
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
}

func (f *ProvincesRequest) Validate() error {
	return nil
}
