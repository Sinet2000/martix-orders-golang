package entity

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
	ID         int64   `json:"id"`
	OrderID    int64   `json:"order_id"`
	ProductID  int64   `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}

type Order struct {
	ID           int64       `json:"id"`
	UserID       int64       `json:"user_id"`
	Items        []OrderItem `json:"items"`
	TotalAmount  float64     `json:"total_amount"`
	Status       OrderStatus `json:"status"`
	ShippingAddr string      `json:"shipping_addr"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
