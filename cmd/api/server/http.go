package server

import (
	"github.com/devrodriguez/first-class-api-go/controllers"
	"github.com/devrodriguez/first-class-api-go/middlewares"
	"github.com/devrodriguez/first-class-api-go/pkg/application"
	"github.com/devrodriguez/first-class-api-go/pkg/infrastructure/mngdb"
	"github.com/devrodriguez/first-class-api-go/pkg/interface/rest"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(client *mongo.Client) *gin.Engine {

	// Check
	chkRepo := mngdb.NewCheckMongoRepo(client)
	chkSrv := application.NewCheckService(chkRepo)
	chkHand := rest.NewCheckHandler(chkSrv)

	// Company
	compRepo := mngdb.NewCompanyMongoRepo(client)
	compSrv := application.NewCompanyService(compRepo)
	compHand := rest.NewCompanyHandler(compSrv)

	// New server
	server := gin.New()

	// Add middlewares
	server.Use(gin.Recovery(), middlewares.Logger())

	// Group api routes
	apiRoutes := server.Group("/api/public")
	{
		apiRoutes.GET("/signin", controllers.SignIn)
	}

	// Endpoints with authentication
	authGroup := server.Group("/api")
	authGroup.Use(middlewares.ValidateAuth())
	{
		authGroup.GET("/companies", compHand.GetAll)
		authGroup.POST("/companies", compHand.Create)
		authGroup.GET("/checks", chkHand.GetChecks)
		authGroup.POST("/checks", chkHand.Create)
		authGroup.PUT("/checks/:id", chkHand.Update)
		authGroup.GET("/login", controllers.Login)
	}

	return server
}
