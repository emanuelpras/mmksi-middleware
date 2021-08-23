package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"middleware-mmksi/dsf/dsf-program/response"
	"net/http"
)

type DsfProgramRepo interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
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
	log.Print("dev: ", url)
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
