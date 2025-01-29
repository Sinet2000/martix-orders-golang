package handler

import (
	"encoding/json"
	"github.com/Sinet2000/Martix-Orders-Go/internal/usecase/order"
	"net/http"

	// "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type OrderHandler struct {
	createOrderUC *order.CreateOrderUseCase
	getOrderUC    *order.GetOrderUseCase
	updateOrderUC *order.UpdateOrderUseCase
}

func NewOrderHandler(
	createOrderUC *order.CreateOrderUseCase,
	getOrderUC *order.GetOrderUseCase,
	updateOrderUC *order.UpdateOrderUseCase,
) *OrderHandler {
	return &OrderHandler{
		createOrderUC: createOrderUC,
		getOrderUC:    getOrderUC,
		updateOrderUC: updateOrderUC,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input order.CreateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := h.createOrderUC.Execute(r.Context(), input)
	if err != nil {
		logger.Error("Failed to create order", zap.Error(err))
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.getOrderUC.Execute(r.Context(), orderID)
	if err != nil {
		logger.Error("Failed to get order", zap.Error(err))
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Router setup
func (h *OrderHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/orders", h.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", h.UpdateOrder).Methods("PUT")
}
