package response

type UnitByModelsResponse struct {
	Message     string         `json:"message"`
	Data        []unitByModels `json:"data"`
	Status      bool           `json:"status"`
	RecordCount int64          `json:"recordCount"`
}

type unitByModels struct {
	ProductModelId int64  `json:"product_model_id"`
	Name           string `json:"name"`
}
