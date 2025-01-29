package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ProductID  primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity   int                `bson:"quantity" json:"quantity"`
	UnitPrice  float64            `bson:"unit_price" json:"unit_price"`
	TotalPrice float64            `bson:"total_price" json:"total_price"`
}

type Order struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Items        []OrderItem        `bson:"items" json:"items"`
	TotalAmount  float64            `bson:"total_amount" json:"total_amount"`
	Status       OrderStatus        `bson:"status" json:"status"`
	ShippingAddr string             `bson:"shipping_addr" json:"shipping_addr"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
