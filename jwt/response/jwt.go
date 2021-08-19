package response

type TokenMmksiResponse struct {
	Token string `json:"access_token"`
}

type TokenRefreshResponse struct {
	TokenRefresh string `json:"refresh_token"`
}
