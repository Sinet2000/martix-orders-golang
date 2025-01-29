package order

import (
	"context"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/entity"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/repository"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
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
	logger.Info("Processing create order request",
		zap.String("user_id", input.UserID.Hex()),
		zap.Int("items_count", len(input.Items)),
	)

	// Business logic validation
	if len(input.Items) == 0 {
		logger.Warn("Attempted to create order with no items",
			zap.String("user_id", input.UserID.Hex()),
		)
		return nil, ErrEmptyOrder
	}

	// Calculate total with detailed logging
	var totalAmount float64
	for _, item := range input.Items {
		totalAmount += item.TotalPrice
		logger.Debug("Processing order item",
			zap.String("product_id", item.ProductID.Hex()),
			zap.Int("quantity", item.Quantity),
			zap.Float64("unit_price", item.UnitPrice),
			zap.Float64("total_price", item.TotalPrice),
		)
	}

	// Create order
	order := &entity.Order{
		UserID:      input.UserID,
		Items:       input.Items,
		TotalAmount: totalAmount,
		Status:      entity.OrderStatusPending,
	}

	if err := uc.orderRepo.Create(ctx, order); err != nil {
		logger.Error("Failed to create order",
			zap.Error(err),
			zap.String("user_id", input.UserID.Hex()),
			zap.Float64("total_amount", totalAmount),
		)
		return nil, err
	}

	logger.Info("Order created successfully",
		zap.String("order_id", order.ID.Hex()),
		zap.String("user_id", order.UserID.Hex()),
		zap.Float64("total_amount", order.TotalAmount),
		zap.String("status", string(order.Status)),
	)

	return order, nil
}
