package response

type AssetCodeResponse struct {
	TargetAssetCode   string `json:"TargetAssetCode"`
	OriginalAssetCode *int64 `json:"OriginalAssetCode"`
}
