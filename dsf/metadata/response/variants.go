package response

type VariantsResponse struct {
	Message     *string   `json:"message"`
	Data        *variants `json:"data"`
	Status      *bool     `json:"status"`
	RecordCount *int      `json:"recordCount"`
}

type variants []struct {
	ProductModelId *int    `json:"product_model_id"`
	Name           *string `json:"name"`
}
