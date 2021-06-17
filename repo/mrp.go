package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/refactory-id/middleware-poc/response"
)

type GetVehiclesParams struct {
	BrandId int64 `json:"brand"`
	ModelId int64 `json:"model"`
}

type GetRegionsParams struct {
	Province string `json:"province"`
}

type GetPredictionParams struct {
	Brand        string `json:"BRAND"`
	Model        string `json:"MODEL"`
	Variant      string `json:"VARIANT"`
	Year         int16  `json:"YEAR"`
	Distance     int64  `json:"DISTANCE"`
	Transmission string `json:"TRANSMISSION"`
	Color        string `json:"COLOR"`
	SellerType   string `json:"TIPE_PENJUAL"`
	City         string `json:"CITY"`
	Province     string `json:"PROVINCE"`
	Company      string `json:"COMPANY"`
}

type MrpRepo interface {
	GetVehicles(params GetVehiclesParams) (*response.GetVehiclesResponse, error)
	GetRegions(params GetRegionsParams) (*response.GetRegionsResponse, error)
	GetPrediction(params GetPredictionParams) (*response.PredictionResponse, error)
}

type mrpRepo struct {
	mrpServer  string
	apiKey     string
	httpClient *http.Client
}

func NewMrpRepo(mrpServer string, apiKey string, httpClient *http.Client) MrpRepo {
	return &mrpRepo{
		mrpServer:  mrpServer,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (r *mrpRepo) GetVehicles(params GetVehiclesParams) (*response.GetVehiclesResponse, error) {
	queryParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/vehicles", r.mrpServer)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(queryParams))
	if err != nil {
		return nil, err
	}

	req.Header.Set("ApiKey", r.apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("mrp: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.GetVehiclesResponse)
	return response, json.Unmarshal(result, response)
}

func (r *mrpRepo) GetRegions(params GetRegionsParams) (*response.GetRegionsResponse, error) {
	queryParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/regions", r.mrpServer)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(queryParams))
	if err != nil {
		return nil, err
	}

	req.Header.Set("ApiKey", r.apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("mrp: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.GetRegionsResponse)
	return response, json.Unmarshal(result, response)

}

func (r *mrpRepo) GetPrediction(params GetPredictionParams) (*response.PredictionResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/prediction", r.mrpServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("ApiKey", r.apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("mrp: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.PredictionResponse)
	return response, json.Unmarshal(result, response)
}
