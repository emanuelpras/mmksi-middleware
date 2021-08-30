package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"middleware-mmksi/dsf/payment/response"
	"net/http"
)

type DsfProgramRepo interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
	GetPackageNames() (*response.GetPackageNames, error)
	GetCarConditions() (*response.GetCarConditions, error)
}

type dsfProgramRepo struct {
	dsfProgramServer string
	httpClient       *http.Client
}

func NewDsfProgramRepo(dsfProgramServer string, httpClient *http.Client) DsfProgramRepo {
	return &dsfProgramRepo{
		dsfProgramServer: dsfProgramServer,
		httpClient:       httpClient,
	}
}

func (r *dsfProgramRepo) GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error) {

	url := fmt.Sprintf("%s/metadata/additional_insurance", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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
		return nil, fmt.Errorf("dsf: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.AdditionalInsuranceResponse)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetPackageNames() (*response.GetPackageNames, error) {

	url := fmt.Sprintf("%s/metadata/packagenames", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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
		return nil, fmt.Errorf("dsf: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.GetPackageNames)
	return response, json.Unmarshal(result, response)
}

func (r *dsfProgramRepo) GetCarConditions() (*response.GetCarConditions, error) {

	url := fmt.Sprintf("%s/metadata/carconditions", r.dsfProgramServer)
	req, err := http.NewRequest("GET", url, nil)
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
		return nil, fmt.Errorf("dsf: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.GetCarConditions)
	return response, json.Unmarshal(result, response)
}
