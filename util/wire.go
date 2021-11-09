package util

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	dsfPaymentRepository "middleware-mmksi/dsf/calculator/repo"
	dsfPaymentService "middleware-mmksi/dsf/calculator/service"
	dsfProgramRepository "middleware-mmksi/dsf/metadata/repo"
	dsfProgramService "middleware-mmksi/dsf/metadata/service"
	mrpController "middleware-mmksi/dsf/mrp/controller"
	mrpRepository "middleware-mmksi/dsf/mrp/repo"
	mrpService "middleware-mmksi/dsf/mrp/service"
	jwtRepository "middleware-mmksi/jwt/repo"
	jwtService "middleware-mmksi/jwt/service"
	mmksiRepository "middleware-mmksi/mmksi/repo"
	mmksiService "middleware-mmksi/mmksi/service"
	salesforceRepository "middleware-mmksi/salesforce/repo"
	salesforceService "middleware-mmksi/salesforce/service"
)

var (
	httpClient *http.Client
)

func ProvideAuthRepo() jwtRepository.JwtRepo {
	return jwtRepository.NewJwtRepo(ProvideHttpClient())
}

func ProvideAuthService() jwtService.JwtService {
	return jwtService.NewJwtService(ProvideAuthRepo())
}

func ProvideMrpController() mrpController.MrpController {
	return mrpController.NewMrpController(ProvideMrpService())
}

func ProvideMrpService() mrpService.MrpService {
	return mrpService.NewMrpService(ProvideMrpRepo())
}

func ProvideMrpRepo() mrpRepository.MrpRepo {
	return mrpRepository.NewMrpRepo(os.Getenv("SERVER_DSF_MRP"), os.Getenv("APIKey_DSF_MRP"), ProvideHttpClient())
}

func ProvideTokenService() mmksiService.MmksiService {
	return mmksiService.NewMmksiService(ProvideTokenRepo())
}

func ProvideTokenRepo() mmksiRepository.MmksiRepo {
	return mmksiRepository.NewMmksiRepo(os.Getenv("SERVER_MMKSI_TOKEN"), ProvideHttpClient())
}

func ProvideMmksiService() mmksiService.MmksiService {
	return mmksiService.NewMmksiService(ProvideMmksiRepo())
}

func ProvideMmksiRepo() mmksiRepository.MmksiRepo {
	return mmksiRepository.NewMmksiRepo(os.Getenv("SERVER_MMKSI_MASTERDATA"), ProvideHttpClient())
}

func ProvideDsfProgramService() dsfProgramService.DsfProgramService {
	return dsfProgramService.NewDsfProgramService(ProvideDsfProgramRepo())
}

func ProvideDsfProgramRepo() dsfProgramRepository.DsfProgramRepo {
	return dsfProgramRepository.NewDsfProgramRepo(os.Getenv("SERVER_DSF_CALCULATOR"), os.Getenv("APIKey_DSF_CALCULATOR"), ProvideHttpClient())
}

func ProvideDsfPaymentService() dsfPaymentService.DsfPaymentService {
	return dsfPaymentService.NewDsfPaymentService(ProvideDsfPaymentRepo())
}

func ProvideDsfPaymentRepo() dsfPaymentRepository.DsfPaymentRepo {
	return dsfPaymentRepository.NewDsfPaymentRepo(os.Getenv("SERVER_DSF_CALCULATOR"), os.Getenv("APIKey_DSF_CALCULATOR"), ProvideHttpClient())
}

func ProvideSalesforceService() salesforceService.SalesforceService {
	return salesforceService.NewSalesforceService(ProvideSalesforceRepo())
}

func ProvideSalesforceRepo() salesforceRepository.SalesforceRepo {
	return salesforceRepository.NewSalesforceRepo(os.Getenv("SERVER_SALESFORCE"), ProvideHttpClient())
}

func ProvideHttpClient() *http.Client {
	transport, ok := http.DefaultTransport.(*http.Transport) // get default roundtripper transport
	if !ok {
		log.Fatal("infra: defaulTransport is not *http.Transport")
	}

	transport.DisableKeepAlives = false
	transport.MaxIdleConns = 256
	transport.MaxIdleConnsPerHost = 256
	transport.MaxConnsPerHost = 0
	transport.ResponseHeaderTimeout = 60 * time.Second
	transport.IdleConnTimeout = 60 * time.Second
	transport.TLSHandshakeTimeout = time.Duration(30) * time.Second
	transport.DialContext = (&net.Dialer{
		Timeout:   time.Duration(60) * time.Second,
		KeepAlive: time.Duration(60) * time.Second,
		DualStack: true,
	}).DialContext

	httpClient = &http.Client{
		Timeout:   time.Duration(60) * time.Second,
		Transport: transport,
	}

	return httpClient
}
