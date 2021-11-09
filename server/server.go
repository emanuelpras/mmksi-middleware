package server

import (
	"log"
	"middleware-mmksi/util"
	"os"

	dsfPaymentControllers "middleware-mmksi/dsf/calculator/controller"
	dsfProgramControllers "middleware-mmksi/dsf/metadata/controller"
	mrpControllers "middleware-mmksi/dsf/mrp/controller"
	jwtControllers "middleware-mmksi/jwt/controller"
	mmksiControllers "middleware-mmksi/mmksi/controller"
	salesforceControllers "middleware-mmksi/salesforce/controller"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "middleware-mmksi/docs"

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

	// @title MMKSI Middleware API Documentation
	// @description MMKSI Middleware API Documentation
	// @version 1.0
	// @host xuu77ziiri.execute-api.us-east-2.amazonaws.com/
	// @basePath development
	// @Schemes  https http

	var (
		authController       jwtControllers.JwtController               = jwtControllers.NewJwtController(util.ProvideAuthService())
		mrpController        mrpControllers.MrpController               = mrpControllers.NewMrpController(util.ProvideMrpService())
		tokenController      mmksiControllers.MmksiController           = mmksiControllers.NewMmksiController(util.ProvideTokenService())
		mmksiController      mmksiControllers.MmksiController           = mmksiControllers.NewMmksiController(util.ProvideMmksiService())
		dsfProgramController dsfProgramControllers.DsfProgramController = dsfProgramControllers.NewDsfProgramController(util.ProvideDsfProgramService())
		dsfPaymentController dsfPaymentControllers.DsfPaymentController = dsfPaymentControllers.NewDsfPaymentController(util.ProvideDsfPaymentService())
		salesforceController salesforceControllers.SalesforceController = salesforceControllers.NewSalesforceController(util.ProvideSalesforceService())
	)

	// Token route
	// Aws cognito sign in
	// you can comment the code if you don't need signin method with aws
	// if you need to signin with aws, you should uncomment this code
	r.POST("/auth/signin", authController.SigninAws)
	r.POST("/auth/resignin", authController.ReSigninAws)

	// Middleware signin method
	// you can comment the code if you want to use middleware signin method
	/* r.POST("/auth/token", authController.CreateToken)
	r.POST("/token/refresh", authController.RefreshToken) */

	// Dsf route
	r.GET("/dsf/tradein/vehicles", mrpController.GetVehicles)
	r.GET("/dsf/tradein/regions", mrpController.GetRegions)
	r.POST("/dsf/tradein/prediction", mrpController.GetPrediction)

	// Metadata route
	r.GET("/dsf/metadata/additionalInsurance", dsfProgramController.GetAdditionalInsurance)
	r.GET("/dsf/metadata/packageNames", dsfProgramController.GetPackageNames)
	r.GET("/dsf/metadata/carConditions", dsfProgramController.GetCarConditions)
	r.POST("/dsf/metadata/packages", dsfProgramController.GetPackages)
	r.GET("/dsf/metadata/variant", dsfProgramController.GetVariants)
	r.GET("/dsf/metadata/paymentTypes", dsfProgramController.GetPaymentTypes)
	r.GET("/dsf/metadata/brands", dsfProgramController.GetBrands)
	r.GET("/dsf/metadata/models", dsfProgramController.GetModels)
	r.GET("/dsf/metadata/vehicleCategory", dsfProgramController.GetVehicleCategory)
	r.GET("/dsf/metadata/branchID", dsfProgramController.GetBranchID)
	r.GET("/dsf/metadata/insuranceTypes", dsfProgramController.GetInsuranceTypes)
	r.GET("/dsf/metadata/insurances", dsfProgramController.GetInsurance)
	r.POST("/dsf/metadata/assetCode", dsfProgramController.GetAssetCode)
	r.GET("/dsf/metadata/provinces", dsfProgramController.GetProvinces)
	r.GET("/dsf/metadata/cities", dsfProgramController.GetCities)

	// Dsf calculator route
	r.POST("/dsf/calculator/perTenor", dsfPaymentController.GetTenor)
	r.POST("/dsf/calculator/allTenors", dsfPaymentController.GetAllTenor)

	// Mmksi master data route
	r.POST("/mmksi/getData", tokenController.GetToken, mmksiController.GetVehicle)
	r.POST("/mmksi/vehicle", tokenController.GetToken, mmksiController.GetVehicleColor)

	// Salesforce route
	r.POST("/salesforce/services/serviceHistory", salesforceController.CheckToken, salesforceController.GetServiceHistory)
	r.POST("/salesforce/services/sparepartSalesHistory", salesforceController.CheckToken, salesforceController.GetSparepartSalesHistory)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
