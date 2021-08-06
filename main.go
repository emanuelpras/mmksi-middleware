package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/refactory-id/middleware-poc/controller"
	"github.com/refactory-id/middleware-poc/util"
)

func main() {
	err := godotenv.Load()

	var (
		jwtController   controller.JwtController   = controller.NewJwtController(util.ProvideJwtService())
		authController  controller.JwtController   = controller.NewJwtController(util.ProvideAuthService())
		mrpController   controller.MrpController   = controller.NewMrpController(util.ProvideMrpService())
		tokenController controller.MmksiController = controller.NewMmksiController(util.ProvideTokenService())
		mmksiController controller.MmksiController = controller.NewMmksiController(util.ProvideMmksiService())
	)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// token route
	r.POST("/auth/token", jwtController.GetFirstToken)

	// dsf route
	r.GET("/dsf/tradein/vehicles", authController.Auth, mrpController.GetVehicles)
	r.GET("/dsf/tradein/regions", authController.Auth, mrpController.GetRegions)
	r.POST("/dsf/tradein/prediction", authController.Auth, mrpController.GetPrediction)

	// mmksi route
	r.POST("/mmksi/getData", authController.Auth, tokenController.GetToken, mmksiController.GetVehicle)
	r.POST("/mmksi/vehicle", authController.Auth, tokenController.GetToken, mmksiController.GetVehicleColor)

	r.Run()
}
