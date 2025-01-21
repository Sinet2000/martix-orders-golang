package route

import (
	"github.com/Sinet2000/Martix-Orders-Go/internal/delivery/http/controller"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderController *controller.OrderController) {
	orderGroup := router.Group("/api/v1/orders")
	orderGroup.Use(middleware.AuthMiddleware())
	{
		orderGroup.POST("/", orderController.CreateOrder)
		orderGroup.GET("/", orderController.ListOrders)
		orderGroup.GET("/:id", orderController.GetOrder)
		orderGroup.PUT("/:id/cancel", orderController.CancelOrder)
		orderGroup.GET("/:id/tracking", orderController.GetOrderTracking)
		orderGroup.POST("/:id/refund", orderController.RequestRefund)
	}
}
