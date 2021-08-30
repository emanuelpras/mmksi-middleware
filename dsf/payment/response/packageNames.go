package response

type GetPackageNames struct {
	Data PackageNames `json:"data"`
}

type PackageNames struct {
	RecordCount int64    `json:"recordCount"`
	PageNumber  int64    `json:"pageNumber"`
	PageSize    int64    `json:"pageSize"`
	Items       []string `json:"items"`
}
