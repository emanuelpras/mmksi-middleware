package response

type TokenOauthResponse struct {
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	ID          string `json:"id"`
	TokenType   string `json:"token_type"`
	Issued_at   string `json:"issued_at"`
	Signature   string `json:"signature"`
}
