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
		mrpController controller.MrpController = controller.NewMrpController(util.ProvideMrpService())
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

	r.Run()
}
