package response

type TokenMmksiResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenAWSResponse struct {
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenAWSResponse struct {
	IDToken   string `json:"id_token"`
	TokenType string `json:"token_type"`
}
