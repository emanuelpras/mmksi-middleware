package repo

import (
	"encoding/json"
	"net/http"

	"middleware-mmksi/jwt/response"
)

type ParamRefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRepo interface {
	RefreshToken(accessToken, refreshToken string) (*response.TokenMmksiResponse, error)
	ValidateRefreshToken(params ParamRefreshToken) (*response.TokenMmksiResponse, error)
}

type refreshTokenRepo struct {
	httpClient *http.Client
}

func NewRefreshTokenRepo(httpClient *http.Client) RefreshTokenRepo {
	return &refreshTokenRepo{
		httpClient: httpClient,
	}
}

func (r *refreshTokenRepo) RefreshToken(accessToken, refreshToken string) (*response.TokenMmksiResponse, error) {

	return &response.TokenMmksiResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *refreshTokenRepo) ValidateRefreshToken(params ParamRefreshToken) (*response.TokenMmksiResponse, error) {

	response := new(response.TokenMmksiResponse)
	return response, json.Unmarshal([]byte(params.RefreshToken), response)
}
