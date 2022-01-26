package server

import (
	"database/sql"
	"log"
	"middleware-mmksi/util"
	"net/http"

	dsfPaymentControllers "middleware-mmksi/dsf/calculator/controller"
	dsfProgramControllers "middleware-mmksi/dsf/metadata/controller"
	mrpControllers "middleware-mmksi/dsf/mrp/controller"
	jwtControllers "middleware-mmksi/jwt/controller"
	mmidControllers "middleware-mmksi/mmid/controller"
	mmksiControllers "middleware-mmksi/mmksi/controller"
	salesforceControllers "middleware-mmksi/salesforce/controller"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "middleware-mmksi/docs"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	DB     *sql.DB
	Router *gin.Engine
	Route  string
}

func NewServer(db *sql.DB, route string) *ApiServer {
	r := gin.New()
	return &ApiServer{
		DB:     db,
		Router: r,
		Route:  route,
	}
}

func (server *ApiServer) ListenAndServe(port string) {
	server.Router.Use(gin.Logger())
	server.Router.Use(gin.Recovery())
	server.registerRoute()

	switch server.Route {
	case "aws-lambda":
		log.Fatal(gateway.ListenAndServe(":"+port, server.Router))
	case "local":
		http.ListenAndServe(":"+port, server.Router)
	default:
		http.ListenAndServe(":"+port, server.Router)
	}
}

func (server *ApiServer) registerRoute() {

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
		mmidController       mmidControllers.MmidController             = mmidControllers.NewMmidController(util.ProvideMmidService())
	)

	// Token route
	// Aws cognito sign in
	// you can comment the code if you don't need signin method with aws
	// if you need to signin with aws, you should uncomment this code
	server.Router.POST("/auth/signin", authController.SigninAws)
	server.Router.POST("/auth/resignin", authController.ReSigninAws)

	// Middleware signin method
	// you can comment the code if you want to use middleware signin method
	/* r.POST("/auth/token", authController.CreateToken)
	r.POST("/token/refresh", authController.RefreshToken) */

	// Dsf route
	server.Router.GET("/dsf/tradein/vehicles", mrpController.GetVehicles)
	server.Router.GET("/dsf/tradein/regions", mrpController.GetRegions)
	server.Router.POST("/dsf/tradein/prediction", mrpController.GetPrediction)

	// Metadata route
	server.Router.GET("/dsf/metadata/additionalInsurance", dsfProgramController.GetAdditionalInsurance)
	server.Router.GET("/dsf/metadata/packageNames", dsfProgramController.GetPackageNames)
	server.Router.GET("/dsf/metadata/carConditions", dsfProgramController.GetCarConditions)
	server.Router.POST("/dsf/metadata/packages", dsfProgramController.GetPackages)
	server.Router.GET("/dsf/metadata/variant", dsfProgramController.GetVariants)
	server.Router.GET("/dsf/metadata/paymentTypes", dsfProgramController.GetPaymentTypes)
	server.Router.GET("/dsf/metadata/brands", dsfProgramController.GetBrands)
	server.Router.GET("/dsf/metadata/models", dsfProgramController.GetModels)
	server.Router.GET("/dsf/metadata/vehicleCategory", dsfProgramController.GetVehicleCategory)
	server.Router.GET("/dsf/metadata/branchID", dsfProgramController.GetBranchID)
	server.Router.GET("/dsf/metadata/insuranceTypes", dsfProgramController.GetInsuranceTypes)
	server.Router.GET("/dsf/metadata/insurances", dsfProgramController.GetInsurance)
	server.Router.POST("/dsf/metadata/assetCode", dsfProgramController.GetAssetCode)
	server.Router.GET("/dsf/metadata/provinces", dsfProgramController.GetProvinces)
	server.Router.GET("/dsf/metadata/cities", dsfProgramController.GetCities)

	// Dsf calculator route
	server.Router.POST("/dsf/calculator/perTenor", dsfPaymentController.GetTenor)
	server.Router.POST("/dsf/calculator/allTenors", dsfPaymentController.GetAllTenor)

	// Mmksi master data route
	server.Router.POST("/mmksi/getData", tokenController.GetToken, mmksiController.GetVehicle)
	server.Router.POST("/mmksi/vehicle", tokenController.GetToken, mmksiController.GetVehicleColor)

	// Salesforce route
	server.Router.POST("/salesforce/services/serviceHistory", salesforceController.CheckToken, salesforceController.GetServiceHistory)
	server.Router.POST("/salesforce/services/sparepartSalesHistory", salesforceController.CheckToken, salesforceController.GetSparepartSalesHistory)

	// Mmid route
	server.Router.POST("/mmid/services/serviceHistory", mmidController.GetServiceHistory)
	server.Router.POST("/mmid/services/serviceHistoryBatch", mmidController.GetServiceHistoryBatch)
	server.Router.POST("/mmid/services/sparepartList", mmidController.GetSparepartList)

	// Swagger route
	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
