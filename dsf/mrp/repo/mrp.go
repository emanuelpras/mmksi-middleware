package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"middleware-mmksi/dsf/mrp/response"
)

type GetVehiclesParams struct {
	BrandName string `form:"brand"`
	ModelName string `form:"model"`
}

type GetRegionsParams struct {
	Province string `form:"province"`
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
	url := fmt.Sprintf("%s/vehicles", r.mrpServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("brand", params.BrandName)
	q.Add("model", params.ModelName)
	req.URL.RawQuery = q.Encode()

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
	url := fmt.Sprintf("%s/regions", r.mrpServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("province", params.Province)
	req.URL.RawQuery = q.Encode()

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

	log.Println("request prediction >>>>> ", req)

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

	log.Println("response prediction >>>>>> ", string(result))

	response := new(response.PredictionResponse)
	return response, json.Unmarshal(result, response)
}
