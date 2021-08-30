package response

type GetPackageNames struct {
	Message  string       `json:"message"`
	Is_Valid bool         `json:"is_valid"`
	Data     PackageNames `json:"data"`
}

type PackageNames struct {
	RecordCount int64    `json:"recordCount"`
	PageNumber  int64    `json:"pageNumber"`
	PageSize    int64    `json:"pageSize"`
	Items       []string `json:"items"`
}
