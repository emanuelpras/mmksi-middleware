package response

type CitiesResponse struct {
	Message          string        `json:"message"`
	Data             []*dataCities `json:"data"`
	Status           bool          `json:"status"`
	CurrentTotalItem int64         `json:"current_total_item"`
	TotalAllItem     int64         `json:"total_all_item"`
}

type dataCities struct {
	ID           int64   `json:"id"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	ProvinceCode string  `json:"province_code"`
	ProvinceName string  `json:"province_name"`
	CreatedDate  *string `json:"created_date"`
	UpdatedDate  *string `json:"updated_date"`
	DeletedDate  *string `json:"deleted_date"`
}
