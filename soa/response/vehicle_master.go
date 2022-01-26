package response

type VehicleMasterResponse struct {
	Meta Meta              `json:"meta"`
	Data VehicleMasterData `json:"data"`
}

type Meta struct {
	Page      int16 `json:"page"`
	Limit     int16 `json:"limit"`
	TotalData int64 `json:"total_data"`
	TotalPage int32 `json:"total_page"`
}

type VehicleMasterData struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	VehicleName  string `json:"vehicle_name"`
	DsfAssetCode string `json:"dsf_asset_code"`
	MmksiType    string `json:"mmksi_type"`
	MmksiColor   string `json:"mmksi_color"`
	Package      string `json:"package"`
	DpMinMax     string `json:"dp_min_max"`
}
