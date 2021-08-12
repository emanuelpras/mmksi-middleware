package main

import (
	"log"
	"os"

	mrpControllers "middleware-mmksi/dsf/mrp/controller"
	jwtControllers "middleware-mmksi/jwt/controller"
	mmksiControllers "middleware-mmksi/mmksi/controller"
	"middleware-mmksi/util"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	RunLocally()
	// RunInAWS()
}

func RunLocally() {
	err := godotenv.Load()

	var (
		authController  jwtControllers.JwtController     = jwtControllers.NewJwtController(util.ProvideAuthService())
		mrpController   mrpControllers.MrpController     = mrpControllers.NewMrpController(util.ProvideMrpService())
		tokenController mmksiControllers.MmksiController = mmksiControllers.NewMmksiController(util.ProvideTokenService())
		mmksiController mmksiControllers.MmksiController = mmksiControllers.NewMmksiController(util.ProvideMmksiService())
	)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// token route
	r.POST("/auth/token", authController.GetFirstToken)

	// dsf route
	r.GET("/dsf/tradein/vehicles", authController.Auth, mrpController.GetVehicles)
	r.GET("/dsf/tradein/regions", authController.Auth, mrpController.GetRegions)
	r.POST("/dsf/tradein/prediction", authController.Auth, mrpController.GetPrediction)

	// mmksi route
	r.POST("/mmksi/getData", authController.Auth, tokenController.GetToken, mmksiController.GetVehicle)
	r.POST("/mmksi/vehicle", authController.Auth, tokenController.GetToken, mmksiController.GetVehicleColor)

	r.Run()
}

func RunInAWS() {
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

	//=====================
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

	return r
}
