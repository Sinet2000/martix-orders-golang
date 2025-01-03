package order

import (
	"database/sql"
	"time"

	"github.com/Sinet2000/Martix-Orders-Go/bootstrap"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(appConfig *bootstrap.AppConfiguration, timeout time.Duration, db *sql.DB, router *gin.RouterGroup) {
	orderRepo := NewOrderRepository(db)
	orderUseCase := NewOrderUseCase(orderRepo, timeout)

	// Initialize controller
	orderController := NewOrderController(orderUseCase)

	// Group all /orders routes
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/", orderController.GetAllOrders) // List all orders
		orderRoutes.GET("/:id", orderController.GetById)   // Get order by ID
		orderRoutes.POST("/", orderController.Create)      // Create a new order
	}
}
