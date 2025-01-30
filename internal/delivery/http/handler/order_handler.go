package handler

import (
	"context"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"github.com/Sinet2000/Martix-Orders-Go/internal/usecase/order"
	"github.com/gin-gonic/gin"
	"net/http"

	// "github.com/gorilla/mux"
	"go.uber.org/zap"
)

type OrderHandler struct {
	createOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderHandler(createOrderUseCase *usecase.CreateOrderUseCase, logger *zap.Logger) *OrderHandler {
	return &OrderHandler{
		createOrderUseCase: createOrderUseCase,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input usecase.CreateOrderInput

	// Decode JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Execute the use case
	orderData, err := h.createOrderUseCase.Execute(context.Background(), input)
	if err != nil {
		logger.Error("Failed to create order", zap.Error(err))

		// Handle specific errors
		switch err {
		case usecase.ErrEmptyOrder:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case usecase.ErrInvalidQuantity:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case usecase.ErrInsufficientStock:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case usecase.ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   orderData,
	})
}

//func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
//	var input order.CreateOrderInput
//	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//		logger.Error("Failed to decode request body", zap.Error(err))
//		http.Error(w, "Invalid request body", http.StatusBadRequest)
//		return
//	}
//
//	order, err := h.createOrderUC.Execute(r.Context(), input)
//	if err != nil {
//		logger.Error("Failed to create order", zap.Error(err))
//		http.Error(w, "Failed to create order", http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(order)
//}
//
//func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	orderID, err := primitive.ObjectIDFromHex(vars["id"])
//	if err != nil {
//		http.Error(w, "Invalid order ID", http.StatusBadRequest)
//		return
//	}
//
//	order, err := h.getOrderUC.Execute(r.Context(), orderID)
//	if err != nil {
//		logger.Error("Failed to get order", zap.Error(err))
//		http.Error(w, "Failed to get order", http.StatusInternalServerError)
//		return
//	}
//
//	if order == nil {
//		http.Error(w, "Order not found", http.StatusNotFound)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(order)
//}

// Router setup
//func (h *OrderHandler) RegisterRoutes(r *mux.Router) {
//	r.HandleFunc("/orders", h.CreateOrder).Methods("POST")
//	r.HandleFunc("/orders/{id}", h.GetOrder).Methods("GET")
//	r.HandleFunc("/orders/{id}", h.UpdateOrder).Methods("PUT")
//}
