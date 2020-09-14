package main

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/devrodriguez/first-class-api-go/cmd/api/server"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	setupLogOutput()
	port := os.Getenv("PORT")

	gin.ForceConsoleColor()

	dbCli := dbConnect()

	server := server.New(dbCli)

	// Serve static files
	server.Static("/s", "./storage")

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

func dbConnect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/locateme?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	// Check connections
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	return client
}
