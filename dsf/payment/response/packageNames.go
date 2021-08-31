package response

type PackageNameResponse struct {
	Message  string       `json:"message"`
	Is_Valid bool         `json:"is_valid"`
	Data     packageNames `json:"data"`
}

type packageNames struct {
	RecordCount int64    `json:"recordCount"`
	PageNumber  int64    `json:"pageNumber"`
	PageSize    int64    `json:"pageSize"`
	Items       []string `json:"items"`
}
