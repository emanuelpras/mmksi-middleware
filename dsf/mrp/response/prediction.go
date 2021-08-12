package response

type PredictionResponse struct {
	PredictionPrice int64   `json:"ProvinceId"`
	AveragePrice    float64 `json:"AveragePrice"`
	MedianPrice     int64   `json:"MedianPrice"`
	MaxPrice        int64   `json:"MaxPrice"`
	MinPrice        int64   `json:"MinPrice"`
}
