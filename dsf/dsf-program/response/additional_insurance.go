package response

type AdditionalInsuranceResponse struct {
	_ []AdditionalInsurance
}

type AdditionalInsurance struct {
	Name string `json:"Name"`
}
