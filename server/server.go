package server

import (
	"log"
	"middleware-mmksi/docs"
	"middleware-mmksi/util"
	"os"

	dsfPaymentControllers "middleware-mmksi/dsf/calculator/controller"
	dsfProgramControllers "middleware-mmksi/dsf/metadata/controller"
	mrpControllers "middleware-mmksi/dsf/mrp/controller"
	jwtControllers "middleware-mmksi/jwt/controller"
	mmksiControllers "middleware-mmksi/mmksi/controller"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewServer(route string) {
	switch route {
	case "aws-lambda":
		awsServer()

	case "local":
	default:
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
		localServer()
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
	docs.SwaggerInfo.Title = "API Documentation"
	docs.SwaggerInfo.Description = "MMKSI Middleware API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	var (
		authController       jwtControllers.JwtController               = jwtControllers.NewJwtController(util.ProvideAuthService())
		mrpController        mrpControllers.MrpController               = mrpControllers.NewMrpController(util.ProvideMrpService())
		tokenController      mmksiControllers.MmksiController           = mmksiControllers.NewMmksiController(util.ProvideTokenService())
		mmksiController      mmksiControllers.MmksiController           = mmksiControllers.NewMmksiController(util.ProvideMmksiService())
		dsfProgramController dsfProgramControllers.DsfProgramController = dsfProgramControllers.NewDsfProgramController(util.ProvideDsfProgramService())
		dsfPaymentController dsfPaymentControllers.DsfPaymentController = dsfPaymentControllers.NewDsfPaymentController(util.ProvideDsfPaymentService())
	)

	// token route
	r.POST("/auth/token", authController.CreateToken)
	r.POST("/token/refresh", authController.RefreshToken)

	// dsf route
	r.GET("/dsf/tradein/vehicles", authController.Auth, mrpController.GetVehicles)
	r.GET("/dsf/tradein/regions", authController.Auth, mrpController.GetRegions)
	r.POST("/dsf/tradein/prediction", authController.Auth, mrpController.GetPrediction)

	// metadata
	r.GET("/dsf/metadata/additionalInsurance", authController.Auth, dsfProgramController.GetAdditionalInsurance)
	r.GET("/dsf/metadata/packageNames", authController.Auth, dsfProgramController.GetPackageNames)
	r.GET("/dsf/metadata/carConditions", authController.Auth, dsfProgramController.GetCarConditions)
	r.POST("/dsf/metadata/packages", authController.Auth, dsfProgramController.GetPackages)
	r.GET("/dsf/metadata/variant", authController.Auth, dsfProgramController.GetVariants)
	r.GET("/dsf/metadata/paymentTypes", authController.Auth, dsfProgramController.GetPaymentTypes)
	r.GET("/dsf/metadata/brands", authController.Auth, dsfProgramController.GetBrands)
	r.GET("/dsf/metadata/models", authController.Auth, dsfProgramController.GetModels)
	r.GET("/dsf/metadata/vehicleCategory", authController.Auth, dsfProgramController.GetVehicleCategory)
	r.GET("/dsf/metadata/branchID", authController.Auth, dsfProgramController.GetBranchID)
	r.GET("/dsf/metadata/insuranceTypes", authController.Auth, dsfProgramController.GetInsuranceTypes)
	r.GET("/dsf/metadata/insurances", authController.Auth, dsfProgramController.GetInsurance)
	r.POST("/dsf/metadata/assetCode", authController.Auth, dsfProgramController.GetAssetCode)
	r.GET("/dsf/metadata/provinces", authController.Auth, dsfProgramController.GetProvinces)
	r.GET("/dsf/metadata/cities", authController.Auth, dsfProgramController.GetCities)

	// calculator
	r.POST("/dsf/calculator/perTenor", authController.Auth, dsfPaymentController.GetTenor)
	r.POST("/dsf/calculator/allTenors", authController.Auth, dsfPaymentController.GetAllTenor)

	// mmksi route
	r.POST("/mmksi/getData", authController.Auth, tokenController.GetToken, mmksiController.GetVehicle)
	r.POST("/mmksi/vehicle", authController.Auth, tokenController.GetToken, mmksiController.GetVehicleColor)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
