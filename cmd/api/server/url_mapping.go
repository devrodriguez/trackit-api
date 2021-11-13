package server

import (
	"github.com/devrodriguez/trackit-go-api/middlewares"
	"github.com/devrodriguez/trackit-go-api/pkg/application"
	"github.com/devrodriguez/trackit-go-api/pkg/infrastructure/db/mysqldb"
	"github.com/devrodriguez/trackit-go-api/pkg/interface/rest"
	"github.com/gin-gonic/gin"
)

func MapURLs(rg *gin.Engine, depend Dependencies) {
	// Auth
	authHand := rest.NewAuthHandler()

	// Check
	chkRepo := mysqldb.NewCheckAdapter(depend.sqlDB)
	chkSrv := application.NewCheckService(chkRepo)
	chkHand := rest.NewCheckHandler(chkSrv)

	// Company
	compRepo := mysqldb.NewCompaniesAdapter(depend.sqlDB)
	compSrv := application.NewCompanyService(compRepo)
	compHand := rest.NewCompanyHandler(compSrv)

	// Group api routes
	apiRoutes := rg.Group("/api/public")
	{
		apiRoutes.GET("/signin", authHand.SignIn)
	}

	// Endpoints with authentication
	authGroup := rg.Group("/api")
	authGroup.Use(middlewares.ValidateAuth())
	{
		authGroup.GET("/companies", compHand.GetAll)
		authGroup.POST("/companies", compHand.Create)
		authGroup.GET("/checks", chkHand.GetChecks)
		authGroup.POST("/checks", chkHand.Create)
		// authGroup.GET("/login", authHand.Login)
	}
}
