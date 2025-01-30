package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// func SetupRouter(orderHandler *handler.OrderHandler) *gin.Engine {
func SetupRouter() *gin.Engine {
	r := gin.New() // or Default
	r.Use(gin.Logger())
	r.Use(gin.Recovery()) // Protect against crashes

	r.GET("/panic", func(c *gin.Context) {
		panic("Something went wrong!")
	})

	// Health Check Route
	r.GET("/api", HealthCheckHandler)

	// Order Routes
	// r.POST("/api/orders", orderHandler.CreateOrder)

	return r
}

func HealthCheckHandler(c *gin.Context) {
	response := gin.H{"message": "Hello, Gin!"}

	c.JSON(http.StatusOK, response)
}
