package main

import (
	"log"
	"time"

	"github.com/Sinet2000/Martix-Orders-Go/bootstrap"
	"github.com/Sinet2000/Martix-Orders-Go/internal/app/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	appConfig := app.AppConfiguration

	db := app.DbClient.Client()
	defer app.CloseDBConnection()

	timeout := time.Duration(appConfig.ContextTimeout) * time.Second

	// Initialize Gin engine
	router := gin.Default()

	// Setup routes
	routes.Setup(appConfig, db, timeout, router)

	// grpc.StartGRPCServer(orderUseCase, appConfig.GRPCAddress)

	// Start the server
	log.Printf("Starting server on %s...\n", appConfig.ServerAddress)
	if err := router.Run(appConfig.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
