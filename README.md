# Mmksi Middleware

This is middleware apps used to get data between DSF and MMKSI.
- DSF can get master data from MMKSI API
- MMKSI can get master data from DSF, payment calculation from DSF and MRP trade in from DSF

This apps will create access token and authentication for both company.

## Prerequisite
1. Go version 1.16.5
2. Run command `go mod tidy` to install dependency

## Project Structure

```
mmksi-middleware
└── dsf
|    ├── controller     # dsf controller directory
|    ├── repo           # dsf repo directory
|    ├── response       # dsf response struct
|    ├── service        # dsf service directory
|        ├── request    # dsf request validation
└── jwt
|    ├── controller     # jwt controller directory
|    ├── repo           # jwt repo directory
|    ├── response       # jwt response struct
|    ├── service        # jwt service directory
|        ├── request    # jwt request validation
└── mmksi
|    ├── controller     # mmksi controller directory
|    ├── repo           # mmksi repo directory
|    ├── response       # mmksi response struct
|    ├── service        # mmksi service directory
|        ├── request    # mmksi request validation
└── server              # routing directory
└── util                # utils file for utilities
main.go                 # main
```

## How to run
- Clone this repository
- To run in local, please run `cp .env.example .env`
- Add your environment variable in file `.env`
- Fill local in `MIDDLEWARE_SERVER` on your `.env`
- Run `go run main.go`
- Run page `localhost:8080/{endpoint}` in your postman

## swagger library `https://github.com/swaggo/gin-swagger`
## How to generate Swagger
- Run command `go mod tidy` to install dependency
- Add `export PATH=$(go env GOPATH)/bin:$PATH` to your bashrc
- Run `swag init` to generate folder docs (docs.go, swagger.json, and swagger.yaml) 
- Run `go run main.go`
- Run page `localhost:8080/swagger/index.html` in your browser
