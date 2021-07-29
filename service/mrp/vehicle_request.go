package mrp

type GetVehicleRequest struct {
	BrandName string `form:"brand"`
	ModelName string `form:"model"`
}

func (f *GetVehicleRequest) Validate() error {
	return nil
}
