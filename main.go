package main

import (
	"io"
	"os"

	"github.com/devrodriguez/first-class-api-go/controllers"
	"github.com/devrodriguez/first-class-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	setupLogOutput()
	port := os.Getenv("PORT")
	server := gin.New()

	gin.ForceConsoleColor()

	// Add middlewares
	server.Use(gin.Recovery(), middlewares.Logger())

	// Group api routes
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/signin", controllers.SignIn)
	}

	authGroup := server.Group("/api/auth")
	authGroup.Use(middlewares.ValidateAuth())
	{
		authGroup.GET("/companies", controllers.GetCompanies)
		authGroup.GET("/checks", controllers.GetChecks)
		authGroup.POST("/checks", controllers.CreateCheck)
		authGroup.PUT("/checks", controllers.UpdateCheck)
		authGroup.GET("/login", controllers.Login)
	}

	if port == "" {
		port = "3001"
	}

	// Run server
	server.Run(":" + port)
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
