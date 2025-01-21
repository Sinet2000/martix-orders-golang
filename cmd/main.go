package main

import (
	"context"
	"fmt"
	"github.com/Sinet2000/Martix-Orders-Go/config"
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
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Errorf("failed to initialize logger %w", err)
			return
		}
	}(logger)

	// Load configuration
	cfg := config.LoadConfig()

	mongoDbContext, err := config.NewMongoService(cfg, logger)
	if err != nil {
		logger.Fatal("Failed to initialize MongoDB", zap.Error(err))
	}
	defer mongoDbContext.Close(context.Background(), logger)

	// Initialize Gin
	gin.SetMode(gin.ReleaseMode)
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
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", zap.Error(err))
	}

	logger.Info("Server exited")
}
