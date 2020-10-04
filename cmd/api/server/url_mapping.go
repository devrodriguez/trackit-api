package server

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/middlewares"
	"github.com/devrodriguez/first-class-api-go/pkg/application"
	"github.com/devrodriguez/first-class-api-go/pkg/infrastructure/persistence/mngdb"

	"github.com/devrodriguez/first-class-api-go/pkg/interface/rest"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MapURLs(rg *gin.Engine, cli *mongo.Client) {

	// Employee
	empRepo := mngdb.NewEmployeeRepository(cli)
	empSrv := application.NewEmployeeService(empRepo)
	empHand := rest.NewEmployeeHandler(empSrv)

	// Auth
	authRepo := mngdb.NewAuthRepository(cli)
	authSrv := application.NewAuthService(authRepo)
	authHand := rest.NewAuthHandler(authSrv)

	// Check
	chkRepo := mngdb.NewCheckMongoRepo(cli)
	chkSrv := application.NewCheckService(chkRepo)
	chkHand := rest.NewCheckHandler(chkSrv)

	// Company
	compRepo := mngdb.NewCompanyMongoRepo(cli)
	compSrv := application.NewCompanyService(compRepo)
	compHand := rest.NewCompanyHandler(compSrv)

	// Worplan
	workRepo := mngdb.NewWorkplanRepository(cli)
	workSrv := application.NewWorkplanService(workRepo)
	workHand := rest.NewWorkplanHandler(workSrv)

	// Group api routes
	apiRoutes := rg.Group("/api/public")
	apiRoutes.Use(middlewares.EnableCORS())
	{
		apiRoutes.GET("/signin", authHand.SignIn)
		apiRoutes.OPTIONS("/signin", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})

		apiRoutes.POST("/register", authHand.Register)
		apiRoutes.OPTIONS("/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
	}

	// Endpoints with authentication
	authGroup := rg.Group("/api")
	authGroup.Use(middlewares.EnableCORS())
	authGroup.Use(middlewares.ValidateAuth())
	{
		authGroup.GET("/companies", compHand.GetAll)
		authGroup.POST("/companies", compHand.Create)
		authGroup.OPTIONS("/companies", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})

		authGroup.GET("/checks", chkHand.GetChecks)
		authGroup.GET("/checks_by", chkHand.GetChecksBy)
		authGroup.OPTIONS("/checks_by", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		authGroup.POST("/checks", chkHand.Create)
		authGroup.PUT("/checks/:id", chkHand.Update)
		authGroup.OPTIONS("/checks", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})

		authGroup.POST("/workplans", workHand.Create)
		authGroup.GET("/workplans/:id", workHand.GetByID)
		authGroup.GET("/employees/:id/workplans", workHand.GetByEmployee)
		authGroup.POST("/employees", empHand.Create)
	}
}
