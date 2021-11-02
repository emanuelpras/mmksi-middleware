package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"middleware-mmksi/salesforce/response"
	"middleware-mmksi/salesforce/service/request"
)

type SalesforceRepo interface {
	GetToken(params request.TokenOauthRequest) (*response.TokenOauthResponse, error)
	GetServiceHistory(params request.ServiceHistoryRequest, header request.HeaderAuthorizationRequest) (*response.ServiceHistoryResponse, error)
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

func (r *salesforceRepo) GetToken(params request.TokenOauthRequest) (*response.TokenOauthResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	log.Print("params", params)

	url := fmt.Sprintf("%s/token", r.salesforceServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "BrowserId=1m7avDt-Eey3klOYloBy9A; CookieConsentPolicy=0:0; LSKey-c$CookieConsentPolicy=0:0")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	log.Print(url, res.Header)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("salesforce: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	log.Print("repo", result)
	if err != nil {
		return nil, err
	}

	response := new(response.TokenOauthResponse)
	return response, json.Unmarshal(result, response)
}

func (r *salesforceRepo) GetServiceHistory(params request.ServiceHistoryRequest, header request.HeaderAuthorizationRequest) (*response.ServiceHistoryResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/services/apexrest/paramServiceHistory", r.salesforceServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// salesforceToken := authorizationSalesforce.TokenType + " " + authorizationSalesforce.AccessToken
	// req.Header.Set("Authorization", "Bearer 00D0l0000000NLq!ARAAQP3pHJ9kPTQduYykvQvPJNbobRzGjsBybPgcN0cAHiK5t2qsLL8zx7Xr0FWe4xZJwR3d8HSWZ1J07J9sldm38tq0v5MZ")
	req.Header.Set("Authorization", header.Authorization)
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

func (r *salesforceRepo) GetSparepartSalesHistory(params request.SparepartSalesHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/services/apexrest/paramSparepartSalesHistory", r.salesforceServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Authorization", "Bearer 00D0l0000000NLq!ARAAQP3pHJ9kPTQduYykvQvPJNbobRzGjsBybPgcN0cAHiK5t2qsLL8zx7Xr0FWe4xZJwR3d8HSWZ1J07J9sldm38tq0v5MZ")
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
