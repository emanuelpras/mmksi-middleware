package request

type BrandsRequest struct {
	Keyword string `form:"keyword"`
	Limit   int64  `form:"limit"`
	Offset  int64  `form:"offset"`
}

func (f *BrandsRequest) Validate() error {
	return nil
}
