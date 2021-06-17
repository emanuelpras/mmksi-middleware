package response

type GetRegionsResponse struct {
	Data []Province
}

type Province struct {
	ProvinceId   int64  `json:"ProvinceId"`
	ProvinceName string `json:"ProvinceName"`
	Cities       []City `json:"Cities"`
}

type City struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
