package response

type CitiesResponse struct {
	Message          *string      `json:"message"`
	Data             *interface{} `json:"data"`
	Status           *bool        `json:"status"`
	CurrentTotalItem *int         `json:"current_total_item"`
	TotalAllItem     *int         `json:"total_all_item"`
}
