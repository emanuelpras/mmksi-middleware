package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"middleware-mmksi/salesforce/response"
	"middleware-mmksi/salesforce/service/request"
)

type SalesforceRepo interface {
	GetTokenSales() (*response.TokenOauthResponse, error)
	GetServiceHistory(params request.ServiceHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error)
	GetSparepartSalesHistory(params request.SparepartSalesHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error)
}

type salesforceRepo struct {
	salesforceServer string
	httpClient       *http.Client
}

func NewSalesforceRepo(salesforceServer string, httpClient *http.Client) SalesforceRepo {
	return &salesforceRepo{
		salesforceServer: salesforceServer,
		httpClient:       httpClient,
	}
}

func (r *salesforceRepo) GetTokenSales() (*response.TokenOauthResponse, error) {

	var p = url.Values{}
	p.Set("grant_type", os.Getenv("SALESFORCE_GRANT_TYPE"))
	p.Set("client_id", os.Getenv("SALESFORCE_CLIENT_ID"))
	p.Set("client_secret", os.Getenv("SALESFORCE_CLIENT_SECRET"))
	p.Set("username", os.Getenv("SALESFORCE_USERNAME"))
	p.Set("password", os.Getenv("SALESFORCE_PASSWORD"))

	payload := bytes.NewBufferString(p.Encode())

	url := fmt.Sprintf("%s/token", os.Getenv("SERVER_SALESFORCE_TOKEN"))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

	response := new(response.TokenOauthResponse)
	return response, json.Unmarshal(result, response)
}

func (r *salesforceRepo) GetServiceHistory(params request.ServiceHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/services/apexrest/paramServiceHistory", authorizationSalesforce.InstanceURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	salesToken := authorizationSalesforce.TokenType + " " + authorizationSalesforce.AccessToken
	req.Header.Set("Authorization", salesToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	os.Setenv("STATUSCODE", strconv.Itoa(res.StatusCode))
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

func (r *salesforceRepo) GetSparepartSalesHistory(params request.SparepartSalesHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/services/apexrest/paramSparepartSalesHistory", authorizationSalesforce.InstanceURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	salesToken := authorizationSalesforce.TokenType + " " + authorizationSalesforce.AccessToken
	req.Header.Set("Authorization", salesToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := r.httpClient.Do(req)
	os.Setenv("STATUSCODE", strconv.Itoa(res.StatusCode))
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
