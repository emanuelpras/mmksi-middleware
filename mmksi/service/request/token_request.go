package request

type TokenRequest struct {
	Clientid   string `form:"clientid"`
	Dealercode string `form:"dealercode"`
	Username   string `form:"username"`
	Password   string `form:"password"`
}
