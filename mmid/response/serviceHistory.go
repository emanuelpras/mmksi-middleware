package response

type ServiceHistoryResponse struct {
	Error  bool   `json:"error"`
	Alerts Alerts `json:"alerts"`
	Data   Data   `json:"data"`
}

type Data struct {
	Status string `json:"status"`
}
type ServiceHistoryBatchResponse struct {
	Error  bool      `json:"error"`
	Alerts Alerts    `json:"alerts"`
	Data   DataBatch `json:"data"`
}

type Alerts struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DataBatch struct {
	Processed    int `json:"processed"`
	Insert_count int `json:"insert_count"`
	Update_count int `json:"update_count"`
}

type SparepartListResponse struct {
	Error  bool      `json:"error"`
	Alerts Alerts    `json:"alerts"`
	Data   DataBatch `json:"data"`
}
