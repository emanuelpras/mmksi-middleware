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
	GetPackageNames(params request.HeaderPackageNameRequest, queryParams request.ParamsPackageNameRequest) (*response.PackageNameResponse, error)
	GetCarConditions() (*response.CarConditionResponse, error)
	GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error)
	GetUnitByModels(paramHeader request.HeaderUnitByModelsRequest) (*response.UnitByModelsResponse, error)
	GetPaymentTypes() (*response.PaymentTypesResponse, error)
	GetBrands(paramHs request.BrandsRequest) (*response.BrandsResponse, error)
	GetVehicleCategory() (*response.VehicleCategory, error)
	GetBranchID() (*response.BranchResponse, error)
	GetInsuranceTypes() (*response.InsuranceTypesResponse, error)
	GetInsurance(params request.InsuranceRequest) (*response.InsuranceResponse, error)
	GetAssetCode(paramHeader request.HeaderAssetCodeRequest, reqBody request.AssetCodeRequest) (*response.AssetCodeResponse, error)
	GetProvinces(params request.ProvincesRequest) (*response.ProvincesResponse, error)
	GetCities(params request.CitiesRequest) (*response.CitiesResponse, error)
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

func (r *dsfProgramRepo) GetPackageNames(params request.HeaderPackageNameRequest, queryParams request.ParamsPackageNameRequest) (*response.PackageNameResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/metadata/"+params.ApplicationName+"/packages/"+params.AssetCode+"/"+params.BranchCode, r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("carCondition", queryParams.CarCondition)
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

func (r *dsfProgramRepo) GetVehicleCategory() (*response.VehicleCategory, error) {
	url := fmt.Sprintf("%s/metadata/vehiclecategories", r.dsfProgramServer)
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

	response := new(response.VehicleCategory)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetBranchID() (*response.BranchResponse, error) {

	url := fmt.Sprintf("%s/metadata/branches", r.dsfProgramServer)
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

	response := new(response.BranchResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetInsuranceTypes() (*response.InsuranceTypesResponse, error) {

	url := fmt.Sprintf("%s/metadata/insuranceTypes", r.dsfProgramServer)
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

	response := new(response.InsuranceTypesResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetInsurance(params request.InsuranceRequest) (*response.InsuranceResponse, error) {

	url := fmt.Sprintf("%s/rates/insurances", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("DsfBranchId", params.DsfBranchId)
	q.Add("VehicleCategory", params.VehicleCategory)
	q.Add("InsuranceTypeCode", params.InsuranceTypeCode)
	q.Add("CarCondition", params.CarCondition)

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

	response := new(response.InsuranceResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetAssetCode(paramHeader request.HeaderAssetCodeRequest, reqBody request.AssetCodeRequest) (*response.AssetCodeResponse, error) {
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/interceptors/"+paramHeader.ApplicationName+"/car", r.dsfProgramServer)
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

	response := new(response.AssetCodeResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetProvinces(params request.ProvincesRequest) (*response.ProvincesResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/location/provinces", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("search", params.Search)
	q.Add("offset", strconv.Itoa(params.Offset))
	q.Add("limit", strconv.Itoa(params.Limit))

	if params.Limit == 0 {
		q.Set("limit", "50")
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

	response := new(response.ProvincesResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetCities(params request.CitiesRequest) (*response.CitiesResponse, error) {

	url := fmt.Sprintf("%s/location/cities", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()

	if params.Search != "" {
		q.Set("search", params.Search)
	}
	if params.ProvinceCode != "" {
		q.Set("province_code", params.ProvinceCode)
	}

	q.Add("offset", strconv.Itoa(params.Offset))
	q.Add("limit", strconv.Itoa(params.Limit))

	if params.Limit == 0 {
		q.Set("limit", "50")
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

	response := new(response.CitiesResponse)

	return response, json.Unmarshal(result, response)
}
