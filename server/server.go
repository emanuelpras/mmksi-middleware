package server

import (
	"log"
	"middleware-mmksi/util"
	"os"

	mrpControllers "middleware-mmksi/dsf/mrp/controller"
	jwtControllers "middleware-mmksi/jwt/controller"
	mmksiControllers "middleware-mmksi/mmksi/controller"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func NewServer(route string) {
	if route == "local" {
		localServer()
	} else if route == "aws" {
		awsServer()
	} else {
		log.Print("route not found, please check your env")
	}
}

func localServer() {
	r := gin.Default()
	registerRoute(r)
	r.Run()
}

func awsServer() {
	addr := ":" + os.Getenv("PORT")
	log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	registerRoute(r)
	return r
}

func registerRoute(r *gin.Engine) {
	var (
		authController  jwtControllers.JwtController     = jwtControllers.NewJwtController(util.ProvideAuthService())
		mrpController   mrpControllers.MrpController     = mrpControllers.NewMrpController(util.ProvideMrpService())
		tokenController mmksiControllers.MmksiController = mmksiControllers.NewMmksiController(util.ProvideTokenService())
		mmksiController mmksiControllers.MmksiController = mmksiControllers.NewMmksiController(util.ProvideMmksiService())
	)

	// token route
	r.POST("/auth/token", authController.GetFirstToken)

	// dsf route
	r.GET("/dsf/tradein/vehicles", authController.Auth, mrpController.GetVehicles)
	r.GET("/dsf/tradein/regions", authController.Auth, mrpController.GetRegions)
	r.POST("/dsf/tradein/prediction", authController.Auth, mrpController.GetPrediction)

	// mmksi route
	r.POST("/mmksi/getData", authController.Auth, tokenController.GetToken, mmksiController.GetVehicle)
	r.POST("/mmksi/vehicle", authController.Auth, tokenController.GetToken, mmksiController.GetVehicleColor)

}
