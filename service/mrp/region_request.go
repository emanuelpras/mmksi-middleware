package mrp

type GetRegionsRequest struct {
	Province string `form:"province"`
}

func (f *GetRegionsRequest) Validate() error {
	return nil
}
