package response

type ProvincesResponse struct {
	Message          string `json:"message"`
	Data             []data
	Status           bool  `json:"status"`
	CurrentTotalItem int64 `json:"current_total_item"`
	TotalAllItem     int64 `json:"total_all_item"`
}

type data struct {
	Code        string  `json:"code"`
	UpdatedDate string  `json:"updated_date"`
	Name        string  `json:"name"`
	Id          int64   `json:"id"`
	DeletedDate *string `json:"deleted_date:"`
	CreatedDate string  `json:"created_date:"`
}
