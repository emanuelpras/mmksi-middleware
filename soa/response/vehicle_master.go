package response

type ListVehicleMasterResponse struct {
	Meta Meta                `json:"meta"`
	Data []VehicleMasterData `json:"data"`
}

type VehicleMasterByAssetCodeResponse struct {
	Data      []VehicleMasterData `json:"data"`
	TotalData int                 `json:"total_data"`
}

type Meta struct {
	Page      int16 `json:"page"`
	Limit     int16 `json:"limit"`
	TotalData int   `json:"total_data"`
	TotalPage int   `json:"total_page"`
}

type VehicleMasterData struct {
	ID           int    `json:"id"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	VehicleName  string `json:"vehicle_name"`
	DsfAssetCode string `json:"dsf_asset_code"`
	MmksiType    string `json:"mmksi_type"`
	MmksiColor   string `json:"mmksi_color"`
	Package      string `json:"package"`
	DpMinMax     string `json:"dp_min_max"`
}
