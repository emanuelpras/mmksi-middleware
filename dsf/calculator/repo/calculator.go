package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"middleware-mmksi/dsf/calculator/response"
	"middleware-mmksi/dsf/calculator/service/request"
	"net/http"
)

type DsfPaymentRepo interface {
	GetTenor(params request.HeaderTenorRequest, reqBody request.TenorRequest) (*response.TenorResponse, error)
	GetAllTenor(params request.HeaderTenorRequest, reqBody request.TenorRequest) (*response.TenorResponse, error)
}

type dsfPaymentRepo struct {
	dsfProgramServer string
	apiKey           string
	httpClient       *http.Client
}

func NewDsfPaymentRepo(dsfProgramServer string, apiKey string, httpClient *http.Client) DsfPaymentRepo {
	return &dsfPaymentRepo{
		dsfProgramServer: dsfProgramServer,
		apiKey:           apiKey,
		httpClient:       httpClient,
	}
}

func (r *dsfPaymentRepo) GetTenor(params request.HeaderTenorRequest, reqBody request.TenorRequest) (*response.TenorResponse, error) {

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/calculator/"+params.ApplicationName, r.dsfProgramServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("ApiKey", r.apiKey)
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("dsf: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.TenorResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfPaymentRepo) GetAllTenor(params request.HeaderTenorRequest, reqBody request.TenorRequest) (*response.TenorResponse, error) {

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/calculator/"+params.ApplicationName+"/alltenors", r.dsfProgramServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("ApiKey", r.apiKey)
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("dsf: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.TenorResponse)
	return response, json.Unmarshal(result, response)
}
