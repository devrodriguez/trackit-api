package server

import (
	"github.com/devrodriguez/first-class-api-go/middlewares"
	"github.com/devrodriguez/first-class-api-go/pkg/application"
	"github.com/devrodriguez/first-class-api-go/pkg/infrastructure/mngdb"
	"github.com/devrodriguez/first-class-api-go/pkg/interface/rest"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MapURLs(rg *gin.Engine, cli *mongo.Client) {
	// Auth
	authHand := rest.NewAuthHandler()

	// Check
	chkRepo := mngdb.NewCheckMongoRepo(cli)
	chkSrv := application.NewCheckService(chkRepo)
	chkHand := rest.NewCheckHandler(chkSrv)

	// Company
	compRepo := mngdb.NewCompanyMongoRepo(cli)
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
		authGroup.PUT("/checks/:id", chkHand.Update)
		// authGroup.GET("/login", authHand.Login)
	}
}
