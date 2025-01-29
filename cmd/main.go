package main

import (
	"context"
	"fmt"
	"github.com/Sinet2000/Martix-Orders-Go/internal/infrastructure/config"
	"github.com/Sinet2000/Martix-Orders-Go/internal/infrastructure/database/mongodb"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// 2. Initialize logger early
	if err := logger.InitLogger(cfg.AppEnv); err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	// Ensure we flush any buffered log entries on shutdown
	defer logger.Sync()

	// 3. Log application startup
	logger.Info("Starting application",
		zap.String("environment", cfg.AppEnv),
		zap.String("port", cfg.Port),
	)

	mongoDbContext, err := mongodb.NewMongoService(&cfg.MongoDB)
	if err != nil {
		logger.Fatal("Failed to initialize MongoDB", zap.Error(err))
	}
	defer mongoDbContext.Close(context.Background())

	// Initialize Gin
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// TODO: Add routes here

	// Create server
	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	// 5. Handle graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	logger.Info("Shutting down application...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", zap.Error(err))
	}

	logger.Info("Server exited")
}
