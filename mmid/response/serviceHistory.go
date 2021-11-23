package response

type ServiceHistoryResponse struct {
	Error  bool   `json:"error"`
	Alerts Alerts `json:"alerts"`
	Data   Data   `json:"data"`
}

type Data interface{}
type ServiceHistoryBatchResponse struct {
	Error  bool   `json:"error"`
	Alerts Alerts `json:"alerts"`
	Data   Data   `json:"data"`
}

type Alerts struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SparepartListResponse struct {
	Error  bool   `json:"error"`
	Alerts Alerts `json:"alerts"`
	Data   Data   `json:"data"`
}
