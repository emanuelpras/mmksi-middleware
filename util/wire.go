package util

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/refactory-id/middleware-poc/controller"
	"github.com/refactory-id/middleware-poc/repo"
	"github.com/refactory-id/middleware-poc/service"
)

var (
	mmksiRepo  repo.MmksiRepo
	mrpRepo    repo.MrpRepo
	httpClient *http.Client
)

func ProvideMrpController() controller.MrpController {
	return controller.NewMrpController(ProvideMrpService())
}

func ProvideMrpService() service.MrpService {
	return service.NewMrpService(ProvideMrpRepo())
}

func ProvideMrpRepo() repo.MrpRepo {
	return repo.NewMrpRepo(os.Getenv("MRP_SERVER"), os.Getenv("MRP_API_KEY"), ProvideHttpClient())
}

func ProvideTokenService() service.MmksiService {
	return service.NewMmksiService(ProvideTokenRepo())
}

func ProvideTokenRepo() repo.MmksiRepo {
	return repo.NewMmksiRepo(os.Getenv("BASEURL_TOKEN"), ProvideHttpClient())
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
