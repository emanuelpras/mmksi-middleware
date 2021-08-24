package repo

import (
	"encoding/json"
	"net/http"

	"middleware-mmksi/jwt/response"
)

type ParamToken struct {
	Company string `json:"company"`
}

type JwtRepo interface {
	CreateToken(accessToken, refreshToken string) (*response.TokenMmksiResponse, error)
	ValidateToken(params ParamToken) (*response.TokenMmksiResponse, error)
}

type jwtRepo struct {
	httpClient *http.Client
}

func NewJwtRepo(httpClient *http.Client) JwtRepo {
	return &jwtRepo{
		httpClient: httpClient,
	}
}

func (r *jwtRepo) CreateToken(accessToken, refreshToken string) (*response.TokenMmksiResponse, error) {

	return &response.TokenMmksiResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *jwtRepo) ValidateToken(params ParamToken) (*response.TokenMmksiResponse, error) {

	response := new(response.TokenMmksiResponse)
	return response, json.Unmarshal([]byte(params.Company), response)
}
