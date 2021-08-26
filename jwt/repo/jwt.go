package repo

import (
	"fmt"
	"net/http"
	"time"

	"middleware-mmksi/jwt/response"
	"middleware-mmksi/jwt/service/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtRepo interface {
	CreateToken(params request.TokenMmksiRequest) (*response.TokenMmksiResponse, error)
	RefreshToken(params request.TokenRefreshRequest) (*response.TokenMmksiResponse, error)
	Auth(gc *gin.Context, params request.AuthRequest) error
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
	token, _ := jwt.Parse(params.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if (claims["company"] == "mmksi") || (claims["company"] == "dsf") {
			company := fmt.Sprintf("%v", claims["company"])

			res, err := r.GenerateToken(company)
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

func (r *jwtRepo) Auth(gc *gin.Context, params request.AuthRequest) error {
	token, _ := jwt.Parse(params.Auth, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if (claims["company"] == "mmksi") || (claims["company"] == "dsf") {
			return nil
		}

		return &response.ErrorResponse{
			ErrorID: 400,
			Msg: map[string]string{
				"en": "Company unregistered",
				"id": "Company tidak terdaftar",
			},
		}
	}
	return &response.ErrorResponse{
		ErrorID: 400,
		Msg: map[string]string{
			"en": "Invalid token or token has expired",
			"id": "Token tidak valid atau token telah kadaluarsa",
		},
	}
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
