package response

import (
	"encoding/json"
)

type InsuranceResponse struct {
	BasicInsuranceRates []basicInsuranceRates
	AdditionalRates     []additionalRates
	LoadingRate         []loadingRate
	PALPRate            json.Number
	TJHRates            []tJHRates
}

type basicInsuranceRates struct {
	MinInsuredAmount json.Number
	MaxInsuredAmount json.Number
	RatePercentage   json.Number
}

type additionalRates struct {
	Name             *string
	RateInPercentage json.Number
}

type loadingRate struct {
	AgeInYears       *int
	RateInPercentage json.Number
}

type tJHRates struct {
	MinInsuredAmount json.Number
	MaxInsuredAmount json.Number
	RatePercentage   json.Number
}
