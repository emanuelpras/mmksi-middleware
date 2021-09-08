package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"middleware-mmksi/dsf/payment/response"
	"middleware-mmksi/dsf/payment/service/request"
	"net/http"
	"strconv"
)

type DsfProgramRepo interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
	GetPackageNames() (*response.PackageNameResponse, error)
	GetCarConditions() (*response.CarConditionResponse, error)
	GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error)
	GetUnitByModels(paramHeader request.HeaderUnitByModelsRequest) (*response.UnitByModelsResponse, error)
	GetPaymentTypes() (*response.PaymentTypesResponse, error)
	GetBrands(paramHs request.BrandsRequest) (*response.BrandsResponse, error)
}

type dsfProgramRepo struct {
	dsfProgramServer string
	apiKey           string
	httpClient       *http.Client
}

func NewDsfProgramRepo(dsfProgramServer string, apiKey string, httpClient *http.Client) DsfProgramRepo {
	return &dsfProgramRepo{
		dsfProgramServer: dsfProgramServer,
		apiKey:           apiKey,
		httpClient:       httpClient,
	}
}

func (r *dsfProgramRepo) GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error) {

	url := fmt.Sprintf("%s/metadata/additional_insurance", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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

	response := new(response.AdditionalInsuranceResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetPackageNames() (*response.PackageNameResponse, error) {

	url := fmt.Sprintf("%s/metadata/packagenames", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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

	response := new(response.PackageNameResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetCarConditions() (*response.CarConditionResponse, error) {

	url := fmt.Sprintf("%s/metadata/carconditions", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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

	response := new(response.CarConditionResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error) {
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/metadata/"+paramHeader.ApplicationName+"/packages", r.dsfProgramServer)
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

	response := new(response.PackageResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetUnitByModels(paramHeader request.HeaderUnitByModelsRequest) (*response.UnitByModelsResponse, error) {

	url := fmt.Sprintf("%s/models/"+paramHeader.ApplicationName+"/units", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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

	response := new(response.UnitByModelsResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetPaymentTypes() (*response.PaymentTypesResponse, error) {

	url := fmt.Sprintf("%s/metadata/paymenttypes", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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

	response := new(response.PaymentTypesResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetBrands(params request.BrandsRequest) (*response.BrandsResponse, error) {

	url := fmt.Sprintf("%s/brands", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("keyword", params.Keyword)
	q.Add("limit", strconv.Itoa(params.Limit))
	q.Add("offset", strconv.Itoa(params.Offset))

	if params.Limit == 0 {
		q.Set("limit", "10")
	}

	req.URL.RawQuery = q.Encode()

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

	response := new(response.BrandsResponse)
	return response, json.Unmarshal(result, response)
}
