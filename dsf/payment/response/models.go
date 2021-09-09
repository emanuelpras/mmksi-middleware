package response

type ModelsResponse struct {
	Message     string   `json:"message"`
	Data        []models `json:"data"`
	Status      bool     `json:"status"`
	RecordCount int64    `json:"recordCount"`
}

type models struct {
	ID           *uint   `json:"id"`
	Model        *string `json:"model"`
	BrandID      *uint   `json:"brand_id"`
	BrandName    *string `json:"brand_name"`
	CategoryID   *uint   `json:"category_id"`
	CategoryName *string `json:"category_name"`
	CategoryCode *string `json:"category_code"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}
