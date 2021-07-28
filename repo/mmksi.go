package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/refactory-id/middleware-poc/response"
)

type GetTokenParams struct {
	Clientid   string `json:"clientid"`
	Dealercode string `json:"dealercode"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type MmksiRepo interface {
	GetToken(params GetTokenParams) (*response.TokenResponse, error)
}

type mmksiRepo struct {
	mmksiServer string
	httpClient  *http.Client
}

func NewMmksiRepo(mmksiServer string, httpClient *http.Client) MmksiRepo {
	return &mmksiRepo{
		mmksiServer: mmksiServer,
		httpClient:  httpClient,
	}
}

func (r *mmksiRepo) GetToken(params GetTokenParams) (*response.TokenResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/token", r.mmksiServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("mmksi: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.TokenResponse)
	return response, json.Unmarshal(result, response)
}
