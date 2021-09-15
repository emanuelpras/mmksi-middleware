package response

type InsuranceResponse struct {
	BasicInsuranceRates *interface{}
	AdditionalRates     *interface{}
	LoadingRate         *interface{}
	PALPRate            *interface{}
	TJHRates            *interface{}
}
