package response

type VehicleResponse struct {
	Total   int64                    `json:"total"`
	Data    []Vehicle                `json:"lst"`
	Message []VehicleMessageResponse `json:"messages"`
}

type Vehicle struct {
	ID              uint   `json:"ID"`
	VehicleType     string `json:"VehicleType"`
	VehicleModel_S1 string `json:"VehicleModel_S1"`
}

type VehicleMessageResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	ErrorCode    int64  `json:"ErrorCode"`
}

type VehicleColorResponse struct {
	Data []VehicleColor `json:"lst"`
}

type VehicleColor struct {
	ID               uint   `json:"ID"`
	VehicleType      string `json:"VehicleType"`
	ColorCode        string `json:"ColorCode"`
	ColorDescription string `json:"ColorDescription"`
	VehicleModel_S1  string `json:"VehicleModel_S1"`
}
