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

	r.POST("/token", func(c *gin.Context) {
		tokenController.GetToken(c)
	})

	r.POST("/mmksi/vehicles", func(c *gin.Context) {
		mmksiController.GetVehicles(c)
	})

	r.Run()
}
