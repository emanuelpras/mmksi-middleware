package response

import "encoding/json"

type ErrorResponse struct {
	ErrorID int               `json:"error_id"`
	Msg     map[string]string `json:"message"`
}

func (c *ErrorResponse) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}
