package usecase

import (
	"context"
	"errors"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/entity"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/repository"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"go.uber.org/zap"
	"time"
)

var (
	ErrEmptyOrder        = errors.New("order must contain at least one item")
	ErrInvalidQuantity   = errors.New("item quantity must be greater than zero")
	ErrInsufficientStock = errors.New("insufficient stock for one or more items")
	ErrProductNotFound   = errors.New("one or more products not found")
)

//type CreateOrderInput struct {
//	UserID int64 `json:"user_id"`
//	Items  []struct {
//		ProductID int64 `json:"product_id"`
//		Quantity  int   `json:"quantity"`
//	} `json:"items"`
//	ShippingAddr string `json:"shipping_addr"`
//}

type CreateOrderInput struct {
	UserID int64 `json:"user_id"`
	Items  []struct {
		ProductID  int64   `json:"product_id"`
		Quantity   int     `json:"quantity"`
		UnitPrice  float64 `json:"unit_price"`
		TotalPrice float64 `json:"total_price"`
	} `json:"items"`
	ShippingAddr string `json:"shipping_addr"`
}

type CreateOrderUseCase struct {
	orderRepo *repository.OrderRepository
	// productRepo *ProductRepository
}

func NewCreateOrderUseCase(
	orderRepo *repository.OrderRepository,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo: orderRepo,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) (*entity.Order, error) {
	logger.Info("Processing create order request",
		zap.Int64("user_id", input.UserID),
		zap.Int("items_count", len(input.Items)),
	)

	if len(input.Items) == 0 {
		logger.Warn("Attempted to create order with no items",
			zap.Int64("user_id", input.UserID),
		)
		return nil, ErrEmptyOrder
	}

	// Create order items and calculate total
	orderItems := make([]entity.OrderItem, 0, len(input.Items))
	var totalAmount float64

	for _, item := range input.Items {
		orderItem := entity.OrderItem{
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: item.TotalPrice,
		}
		orderItems = append(orderItems, orderItem)
		totalAmount += item.TotalPrice

		logger.Debug("Processing order item",
			zap.Int64("product_id", item.ProductID),
			zap.Int("quantity", item.Quantity),
			zap.Float64("unit_price", item.UnitPrice),
			zap.Float64("total_price", item.TotalPrice),
		)
	}

	order := &entity.Order{
		UserID:       input.UserID,
		Items:        orderItems,
		TotalAmount:  totalAmount,
		Status:       entity.OrderStatusPending,
		ShippingAddr: input.ShippingAddr,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := uc.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		logger.Error("Failed to create order",
			zap.Error(err),
			zap.Int64("user_id", input.UserID),
			zap.Float64("total_amount", totalAmount),
		)
		return nil, err
	}

	logger.Info("Order created successfully",
		zap.Int64("order_id", order.ID),
		zap.Int64("user_id", order.UserID),
		zap.Float64("total_amount", order.TotalAmount),
		zap.String("status", string(order.Status)),
	)

	return order, nil
}

//type CreateOrderInput struct {
//	UserID       primitive.ObjectID `json:"user_id"`
//	Items        []entity.OrderItem `json:"items"`
//	ShippingAddr string             `json:"shipping_addr"`
//}
//
//type CreateOrderUseCase struct {
//	orderRepo repository.OrderRepository
//}
//
//func NewCreateOrderUseCase(orderRepo repository.OrderRepository) *CreateOrderUseCase {
//	return &CreateOrderUseCase{
//		orderRepo: orderRepo,
//	}
//}
//
//func (uc *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) (*entity.Order, error) {
//	logger.Info("Processing create order request",
//		zap.String("user_id", input.UserID.Hex()),
//		zap.Int("items_count", len(input.Items)),
//	)
//
//	// Business logic validation
//	if len(input.Items) == 0 {
//		logger.Warn("Attempted to create order with no items",
//			zap.String("user_id", input.UserID.Hex()),
//		)
//		return nil, ErrEmptyOrder
//	}
//
//	// Calculate total with detailed logging
//	var totalAmount float64
//	for _, item := range input.Items {
//		totalAmount += item.TotalPrice
//		logger.Debug("Processing order item",
//			zap.String("product_id", item.ProductID.Hex()),
//			zap.Int("quantity", item.Quantity),
//			zap.Float64("unit_price", item.UnitPrice),
//			zap.Float64("total_price", item.TotalPrice),
//		)
//	}
//
//	// Create order
//	order := &entity.Order{
//		UserID:      input.UserID,
//		Items:       input.Items,
//		TotalAmount: totalAmount,
//		Status:      entity.OrderStatusPending,
//	}
//
//	if err := uc.orderRepo.Create(ctx, order); err != nil {
//		logger.Error("Failed to create order",
//			zap.Error(err),
//			zap.String("user_id", input.UserID.Hex()),
//			zap.Float64("total_amount", totalAmount),
//		)
//		return nil, err
//	}
//
//	logger.Info("Order created successfully",
//		zap.String("order_id", order.ID.Hex()),
//		zap.String("user_id", order.UserID.Hex()),
//		zap.Float64("total_amount", order.TotalAmount),
//		zap.String("status", string(order.Status)),
//	)
//
//	return order, nil
//}
