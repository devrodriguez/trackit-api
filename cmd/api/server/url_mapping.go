package server

import (
	"github.com/devrodriguez/trackit-go-api/cmd/api/server/handlers"
	"github.com/devrodriguez/trackit-go-api/cmd/api/server/middlewares"
	"github.com/devrodriguez/trackit-go-api/pkg/application"
	"github.com/devrodriguez/trackit-go-api/pkg/infrastructure/db/mysqldb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapURLs(rg *gin.Engine, depend Dependencies) {

	// Check
	chkRepo := mysqldb.NewCheckAdapter(depend.sqlDB)
	chkSrv := application.NewCheckService(chkRepo)
	chkHand := handlers.NewCheckHandler(chkSrv)

	// Company
	compRepo := mysqldb.NewCompaniesAdapter(depend.sqlDB)
	compSrv := application.NewCompanyService(compRepo)
	compHand := handlers.NewCompanyHandler(compSrv)

	// Employee
	empRepo := mysqldb.NewEmployeeAdapter(depend.sqlDB)
	empSrv := application.NewEmployeeSrv(empRepo)
	empHand := handlers.NewEmployeeHandler(empSrv)

	// User
	usrRepo := mysqldb.NewUserAdapter(depend.sqlDB)
	usrSrv := application.NewUserSrv(usrRepo)
	usrHand := handlers.NewUserHandler(usrSrv)

	// Auth
	authHand := handlers.NewAuthHandler(usrRepo)

	// Group api routes
	apiRoutes := rg.Group("/api/public")
	apiRoutes.Use(middlewares.EnableCORS())
	{
		apiRoutes.GET("/signin", authHand.SignIn)
		apiRoutes.OPTIONS("/signin", func(c *gin.Context) { c.JSON(http.StatusOK, nil)	})
	}

	// Endpoints with authentication
	authGroup := rg.Group("/api")
	authGroup.Use(middlewares.EnableCORS(), middlewares.ValidateAuth())
	{
		authGroup.GET("/companies", compHand.GetAll)
		authGroup.POST("/companies", compHand.Create)
		authGroup.GET("/employees/:emp_id/checks", chkHand.GetChecks)
		authGroup.OPTIONS("/employees/:emp_id/checks", func(c *gin.Context) { c.JSON(http.StatusOK, nil)	})
		authGroup.POST("/checks", chkHand.Create)
		authGroup.POST("/employees", empHand.Create)
		authGroup.POST("/users", usrHand.Create)
		authGroup.GET("/login", authHand.Login)
		authGroup.OPTIONS("/login", func(c *gin.Context) { c.JSON(http.StatusOK, nil)	})
	}
}
