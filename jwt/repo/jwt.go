package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"middleware-mmksi/jwt/response"
)

type ParamToken struct {
	Company string `json:"company"`
}

type ParamRefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type JwtRepo interface {
	CreateToken(params ParamToken) (*response.TokenMmksiResponse, error)
	RefreshToken(params ParamRefreshToken) (*response.TokenRefreshResponse, error)
}

type jwtRepo struct {
	jwtServer  string
	httpClient *http.Client
}

func NewJwtRepo(httpClient *http.Client) JwtRepo {
	return &jwtRepo{
		httpClient: httpClient,
	}
}

func (r *jwtRepo) CreateToken(params ParamToken) (*response.TokenMmksiResponse, error) {
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

	response := new(response.TokenMmksiResponse)
	return response, json.Unmarshal(result, response)
}

func (r *jwtRepo) RefreshToken(params ParamRefreshToken) (*response.TokenRefreshResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s", r.jwtServer)
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

	response := new(response.TokenRefreshResponse)
	return response, json.Unmarshal(result, response)
}
