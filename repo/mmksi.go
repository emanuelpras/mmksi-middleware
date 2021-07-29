package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	response "github.com/refactory-id/middleware-poc/response/mmksi"
)

type GetTokenParams struct {
	Clientid   string `json:"clientid"`
	Dealercode string `json:"dealercode"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type GetVehicleParams struct {
	Page int64 `json:"page"`
}

type GetVehicleHeaderAuthorization struct {
	Authorization string `json:"Authorization"`
}

type MmksiRepo interface {
	GetToken(params GetTokenParams) (*response.TokenResponse, error)
	GetVehicles(params GetVehicleParams, paramHeader GetVehicleHeaderAuthorization) (*response.VehicleResponse, error)
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

func (r *mmksiRepo) GetVehicles(params GetVehicleParams, authorizationMmksi GetVehicleHeaderAuthorization) (*response.VehicleResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/Master/WebsiteVehicleType/Read", r.mmksiServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authorizationMmksi.Authorization)
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

	response := new(response.VehicleResponse)
	return response, json.Unmarshal(result, response)
}
