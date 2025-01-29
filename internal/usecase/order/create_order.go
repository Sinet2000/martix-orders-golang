package order

import (
	"context"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/entity"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type CreateOrderInput struct {
	UserID       primitive.ObjectID `json:"user_id"`
	Items        []entity.OrderItem `json:"items"`
	ShippingAddr string             `json:"shipping_addr"`
}

type CreateOrderUseCase struct {
	orderRepo repository.OrderRepository
}

func NewCreateOrderUseCase(orderRepo repository.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo: orderRepo,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) (*entity.Order, error) {
	totalAmount := 0.0
	for _, item := range input.Items {
		totalAmount += item.TotalPrice
	}

	order := &entity.Order{
		UserID:       input.UserID,
		Items:        input.Items,
		TotalAmount:  totalAmount,
		Status:       entity.OrderStatusPending,
		ShippingAddr: input.ShippingAddr,
	}

	if err := uc.orderRepo.Create(ctx, order); err != nil {
		logger.Error("Failed to create order", zap.Error(err))
		return nil, err
	}

	return order, nil
}
