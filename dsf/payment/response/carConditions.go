package response

type GetCarConditions []struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
