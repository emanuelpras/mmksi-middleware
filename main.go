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
	r.GET("/vehicles", func(c *gin.Context) {
		mrpController.GetVehicles(c)
	})

	r.GET("/regions", func(c *gin.Context) {
		mrpController.GetRegions(c)
	})

	r.POST("/prediction", func(c *gin.Context) {
		mrpController.GetPrediction(c)
	})

	r.POST("/create/token", func(c *gin.Context) {
		jwtController.GetFirstToken(c)
	})

	r.POST("/getData",
		func(c *gin.Context) {
			authController.Auth(c)
		},
		func(c *gin.Context) {
			tokenController.GetToken(c)
		},
		func(c *gin.Context) {
			mmksiController.GetVehicle(c)
		},
	)

	r.Run()
}
