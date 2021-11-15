package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"middleware-mmksi/mmid/response"
	"middleware-mmksi/mmid/service/request"
)

type MmidRepo interface {
	GetServiceHistory(params request.ServiceHistoryRequest) (*response.ServiceHistoryResponse, error)
	GetServiceHistoryBatch(params request.Batch) (*response.ServiceHistoryBatchResponse, error)
}

type mmidRepo struct {
	mmidServer string
	httpClient *http.Client
}

func NewMmidRepo(mmidServer string, httpClient *http.Client) MmidRepo {
	return &mmidRepo{
		mmidServer: mmidServer,
		httpClient: httpClient,
	}
}

func (r *mmidRepo) GetServiceHistory(params request.ServiceHistoryRequest) (*response.ServiceHistoryResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/service-history/data", r.mmidServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Token", os.Getenv("TOKEN_MMID"))
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("salesforce: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.ServiceHistoryResponse)
	return response, json.Unmarshal(result, response)
}

func (r *mmidRepo) GetServiceHistoryBatch(params request.Batch) (*response.ServiceHistoryBatchResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/service-history/data-batch", r.mmidServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Token", os.Getenv("TOKEN_MMID"))
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("salesforce: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.ServiceHistoryBatchResponse)
	return response, json.Unmarshal(result, response)
}
