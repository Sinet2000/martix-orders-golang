package controller

import (
	"github.com/Sinet2000/Martix-Orders-Go/internal/entity"
	"github.com/Sinet2000/Martix-Orders-Go/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type OrderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) *OrderController {
	return &OrderController{
		orderUseCase: orderUseCase,
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order entity.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.orderUseCase.CreateOrder(ctx, &order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) CancelOrder(ctx *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	if err := c.orderUseCase.CancelOrder(ctx, orderID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
