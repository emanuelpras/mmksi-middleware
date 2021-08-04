package mrp

type GetVehiclesResponse struct {
	Data []Vehicle
}

type Vehicle struct {
	BrandId   int64          `json:"BrandId"`
	BrandName string         `json:"BrandName"`
	Models    []VehicleModel `json:"Models"`
}

type VehicleModel struct {
	ModelId   int64     `json:"ModelId"`
	ModelName string    `json:"ModelName"`
	Variants  []Variant `json:"Variants"`
}

type Variant struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
