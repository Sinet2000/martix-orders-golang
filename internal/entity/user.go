package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Email          string               `bson:"email" json:"email" validate:"required,email"`
	Password       string               `bson:"password" json:"-" validate:"required,min=8"`
	FirstName      string               `bson:"firstName" json:"firstName" validate:"required"`
	LastName       string               `bson:"lastName" json:"lastName" validate:"required"`
	PhoneNumber    string               `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Type           string               `bson:"type" json:"type" validate:"oneof=regular premium business"` // User type for discounts
	Status         string               `bson:"status" json:"status" validate:"oneof=active inactive blocked"`
	Addresses      []Address            `bson:"addresses" json:"addresses"`
	PaymentMethods []PaymentMethod      `bson:"paymentMethods" json:"paymentMethods"`
	Orders         []primitive.ObjectID `bson:"orders" json:"orders"`
	LoyaltyPoints  int                  `bson:"loyaltyPoints" json:"loyaltyPoints"`
	LastLoginAt    time.Time            `bson:"lastLoginAt" json:"lastLoginAt"`
	CreatedAt      time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt" json:"updatedAt"`
}

type Address struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type       string             `bson:"type" json:"type" validate:"oneof=home work other"`
	IsDefault  bool               `bson:"isDefault" json:"isDefault"`
	Street     string             `bson:"street" json:"street" validate:"required"`
	City       string             `bson:"city" json:"city" validate:"required"`
	State      string             `bson:"state" json:"state" validate:"required"`
	Country    string             `bson:"country" json:"country" validate:"required"`
	PostalCode string             `bson:"postalCode" json:"postalCode" validate:"required"`
}

type PaymentMethod struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type           string             `bson:"type" json:"type" validate:"oneof=credit debit paypal"`
	IsDefault      bool               `bson:"isDefault" json:"isDefault"`
	Provider       string             `bson:"provider" json:"provider"`
	LastFourDigits string             `bson:"lastFourDigits" json:"lastFourDigits"`
	ExpiryMonth    int                `bson:"expiryMonth" json:"expiryMonth"`
	ExpiryYear     int                `bson:"expiryYear" json:"expiryYear"`
	HolderName     string             `bson:"holderName" json:"holderName"`
}
