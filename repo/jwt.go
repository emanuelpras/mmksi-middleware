package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	response "github.com/refactory-id/middleware-poc/response/jwt"
)

type ParamToken struct {
	Company string `json:"company"`
}

type JwtRepo interface {
	GetFirstToken(params ParamToken) (*response.FirtsTokenResponse, error)
}

type jwtRepo struct {
	jwtServer  string
	httpClient *http.Client
}

func NewJwtRepo(jwtServer string, httpClient *http.Client) JwtRepo {
	return &jwtRepo{
		jwtServer:  jwtServer,
		httpClient: httpClient,
	}
}

func (r *jwtRepo) GetFirstToken(params ParamToken) (*response.FirtsTokenResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/create/token", r.jwtServer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
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
		return nil, fmt.Errorf("jwt: response status %d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(response.FirtsTokenResponse)
	return response, json.Unmarshal(result, response)
}
