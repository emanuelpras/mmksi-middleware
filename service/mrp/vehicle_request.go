package mrp

type GetVehicleRequest struct {
	BrandId string `form:"brand"`
	ModelId string `form:"model"`
}

func (f *GetVehicleRequest) Validate() error {
	return nil
}
