package response

type TenorResponse struct {
	Message    *string     `json:"message"`
	Data       interface{} `json:"data"`
	Status     *bool       `json:"status"`
	StatusCode *string     `json:"statusCode"`
}

type AllTenorResponse struct {
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
	Status  *bool       `json:"status"`
}
