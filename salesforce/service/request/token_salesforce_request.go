package request

type TokenOauthRequest struct {
	GrantType    string `form:"grant_type"`
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Username     string `form:"username"`
	Password     string `form:"password"`
}
