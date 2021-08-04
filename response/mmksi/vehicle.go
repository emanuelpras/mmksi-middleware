package mmksi

type VehicleResponse struct {
	Data []Vehicle `json:"lst"`
}

type Vehicle struct {
	ID              uint   `json:"ID"`
	VehicleType     string `json:"VehicleType"`
	VehicleModel_S1 string `json:"VehicleModel_S1"`
}
