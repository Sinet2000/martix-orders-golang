package routes

import (
	"database/sql"
	"time"

	"github.com/Sinet2000/Martix-Orders-Go/bootstrap"
	"github.com/Sinet2000/Martix-Orders-Go/internal/app/order"
	"github.com/gin-gonic/gin"
)

// Setup initializes all routes for the application
func Setup(appConfig *bootstrap.AppConfiguration, db *sql.DB, timeout time.Duration, router *gin.Engine) {
	// API version grouping
	api := router.Group("/api/v1")

	// Apply global filters
	// router.Use(filters.PanicRecovery())

	// // Apply global middlewares
	// router.Use(middleware.AuthMiddleware())

	// Order module routes
	order.RegisterRoutes(appConfig, timeout, db, api)

	// Add more modules here in the future, e.g., Product, User
}
