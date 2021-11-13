package main

import (
	"io"
	"os"

	"github.com/devrodriguez/trackit-go-api/cmd/api/server"
	"github.com/gin-gonic/gin"
)

func main() {

	setupLogOutput()

	gin.ForceConsoleColor()

	server := server.New()

	port := os.Getenv("PORT")

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
