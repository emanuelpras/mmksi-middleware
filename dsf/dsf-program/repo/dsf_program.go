package repo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AdditionalResponse struct {
	_ additionalResponse
}
type additionalResponse struct {
	Name string `json:"Name"`
}

type DsfProgramRepo interface {
	GetAdditionalInsurance() (*AdditionalResponse, error)
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

func (r *dsfProgramRepo) GetAdditionalInsurance() (*AdditionalResponse, error) {

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

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Print("body", res.Body)

	response := new(AdditionalResponse)
	log.Print("res repo", response, err)
	return response, err
}
