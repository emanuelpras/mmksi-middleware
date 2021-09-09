package request

type BrandsRequest struct {
	Keyword string `form:"keyword"`
	Limit   int    `form:"limit"`
	Offset  int    `form:"offset"`
}

func (f *BrandsRequest) Validate() error {
	return nil
}
