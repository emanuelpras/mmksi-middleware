package request

type ModelsRequest struct {
	Brand string `form:"brand"`
}

func (f *ModelsRequest) Validate() error {
	return nil
}
