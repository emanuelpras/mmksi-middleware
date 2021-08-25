package repo

import (
	"net/http"
	"time"

	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
)

type JwtRepo interface {
	CreateToken(params request.TokenMmksiRequest) (*response.TokenMmksiResponse, error)
	RefreshToken(params request.TokenRefreshRequest) (*response.TokenMmksiResponse, error)
}

type jwtRepo struct {
	httpClient *http.Client
}

func NewJwtRepo(httpClient *http.Client) JwtRepo {
	return &jwtRepo{
		httpClient: httpClient,
	}
}

func (r *jwtRepo) CreateToken(params request.TokenMmksiRequest) (*response.TokenMmksiResponse, error) {
	return r.GenerateToken(params.Company)
}

func (r *jwtRepo) RefreshToken(params request.TokenRefreshRequest) (*response.TokenMmksiResponse, error) {
	return r.GenerateToken(params.RefreshToken)
}

func (r *jwtRepo) GenerateToken(params string) (*response.TokenMmksiResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["company"] = params
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return &response.TokenMmksiResponse{}, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["company"] = params
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	refreshToken, err2 := refresh.SignedString([]byte("secret"))
	if err2 != nil {
		return &response.TokenMmksiResponse{}, err2
	}

	return &response.TokenMmksiResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
