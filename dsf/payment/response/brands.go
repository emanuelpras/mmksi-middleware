package response

type BrandsResponse struct {
	Message     string    `json:"message"`
	Data        []*brands `json:"data"`
	Status      bool      `json:"status"`
	RecordCount int64     `json:"recordCount"`
}

type brands struct {
	DeletedAt     *string `json:"deleted_at"`
	Createdat     *string `json:"created_at"`
	BrandImageUrl *string `json:"brand_image_url"`
	ID            int64   `json:"id"`
	IsVehicle     bool    `json:"is_vehicle"`
	UpdatedAt     *string `json:"updated_at"`
	Name          *string `json:"name"`
}
