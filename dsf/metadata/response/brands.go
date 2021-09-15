package response

type BrandsResponse struct {
	Message     *string      `json:"message"`
	Data        *interface{} `json:"data"`
	Status      *bool        `json:"status"`
	RecordCount *int         `json:"recordCount"`
}
