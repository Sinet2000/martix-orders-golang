package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
	Status      string             `bson:"status" json:"status"`
	TotalAmount float64            `bson:"totalAmount" json:"totalAmount"`
	Items       []OrderItem        `bson:"items" json:"items"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type OrderItem struct {
	ProductID primitive.ObjectID `bson:"productId" json:"productId"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"price" json:"price"`
	Subtotal  float64            `bson:"subtotal" json:"subtotal"`
}
