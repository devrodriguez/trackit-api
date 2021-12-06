package server

import (
	"github.com/devrodriguez/trackit-go-api/cmd/api/server/middlewares"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	// New server
	server := gin.New()

	// Serve static files
	server.Static("/s", "./storage")

	// Add middlewares
	server.Use(gin.Recovery(), middlewares.Logger())
	// server.Use(middlewares.CacheControl())

	// Build dependencies
	depend := BuildDependencies()

	// Map URLs
	MapURLs(server, depend)

	return server
}
