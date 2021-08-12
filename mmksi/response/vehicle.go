package response

type VehicleResponse struct {
	Data []Vehicle `json:"lst"`
}

type Vehicle struct {
	ID              uint   `json:"ID"`
	VehicleType     string `json:"VehicleType"`
	VehicleModel_S1 string `json:"VehicleModel_S1"`
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
