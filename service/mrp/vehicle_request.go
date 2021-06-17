package mrp

type GetVehicleRequest struct {
	BrandId int64 `form:"brand"`
	ModelId int64 `form:"model"`
}

func (f *GetVehicleRequest) Validate() error {
	return nil
}
