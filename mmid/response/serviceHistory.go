package response

type ServiceHistoryResponse struct {
	Error  bool        `json:"error"`
	Alerts Alerts      `json:"alerts"`
	Data   interface{} `json:"data"`
}

type Alerts struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
