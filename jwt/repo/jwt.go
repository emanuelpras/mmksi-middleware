package repo

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Timeout struct {
	TimeoutAccessToken  string
	TimeoutRefreshToken string
}
type JwtRepo interface {
	CreateToken(params request.TokenMmksiRequest, timeout Timeout) (*response.TokenMmksiResponse, error)
	RefreshToken(params request.TokenRefreshRequest, timeout Timeout) (*response.TokenMmksiResponse, error)
	Auth(gc *gin.Context, params request.AuthRequest, timeout Timeout) error
}
type jwtRepo struct {
	httpClient *http.Client
}

func NewJwtRepo(httpClient *http.Client) JwtRepo {
	return &jwtRepo{
		httpClient: httpClient,
	}
}

func (r *jwtRepo) CreateToken(params request.TokenMmksiRequest, timeout Timeout) (*response.TokenMmksiResponse, error) {

	return r.GenerateToken(params.Company, timeout)

}

func (r *jwtRepo) RefreshToken(params request.TokenRefreshRequest, timeout Timeout) (*response.TokenMmksiResponse, error) {

	return r.TokenValidation(params.RefreshToken, timeout)

}

func (r *jwtRepo) Auth(gc *gin.Context, params request.AuthRequest, timeout Timeout) error {

	_, err := r.TokenValidation(params.Auth, timeout)
	if err != nil {
		return err
	}
	return nil

}

func (r *jwtRepo) TokenValidation(params string, timeout Timeout) (*response.TokenMmksiResponse, error) {

	token, _ := jwt.Parse(params, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if (claims["company"] == "mmksi") || (claims["company"] == "dsf") {
			company := fmt.Sprintf("%v", claims["company"])

			res, err := r.GenerateToken(company, timeout)
			if err != nil {
				return nil, err
			}
			return res, err
		}

		return nil, &response.ErrorResponse{
			ErrorID: 400,
			Msg: map[string]string{
				"en": "Company unregistered",
				"id": "Company tidak terdaftar",
			},
		}
	}

	return nil, &response.ErrorResponse{
		ErrorID: 400,
		Msg: map[string]string{
			"en": "Invalid token or token has expired",
			"id": "Token tidak valid atau token telah kadaluarsa",
		},
	}

}

func (r *jwtRepo) GenerateToken(company string, timeout Timeout) (*response.TokenMmksiResponse, error) {

	timeoutAccessToken, _ := strconv.Atoi(timeout.TimeoutAccessToken)
	timeoutRefreshToken, _ := strconv.Atoi(timeout.TimeoutRefreshToken)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["company"] = company
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeoutAccessToken)).Unix()
	accessToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return &response.TokenMmksiResponse{}, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["company"] = company
	rtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(timeoutRefreshToken)).Unix()
	refreshToken, err2 := refresh.SignedString([]byte("secret"))

	if err2 != nil {
		return &response.TokenMmksiResponse{}, err2
	}

	return &response.TokenMmksiResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
