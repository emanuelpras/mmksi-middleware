package response

type PaymentTypesResponse []struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
