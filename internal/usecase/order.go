package usecase

import (
	"context"
	"errors"
	"github.com/Sinet2000/Martix-Orders-Go/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"time"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Order, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	GetUserOrders(ctx context.Context, userID primitive.ObjectID) ([]entity.Order, error)
}

type OrderUseCase struct {
	repo   OrderRepository
	logger *zap.Logger
	//paymentService  PaymentService
	//shippingService ShippingService
	//discountService DiscountService
	//eventPublisher  EventPublisher
}

func NewOrderUseCase(repo OrderRepository, logger *zap.Logger) *OrderUseCase {
	return &OrderUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order *entity.Order) error {
	if len(order.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	var totalAmount float64
	for i := range order.Items {
		order.Items[i].Subtotal = order.Items[i].Price * float64(order.Items[i].Quantity)
		totalAmount += order.Items[i].Subtotal
	}

	order.TotalAmount = totalAmount
	order.Status = "pending"

	return uc.repo.Create(ctx, order)
}

func (uc *OrderUseCase) GetOrder(ctx context.Context, id primitive.ObjectID) (*entity.Order, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *OrderUseCase) UpdateOrderStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	validStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"shipped":    true,
		"delivered":  true,
		"cancelled":  true,
	}

	if !validStatuses[status] {
		return errors.New("invalid order status")
	}

	return uc.repo.UpdateStatus(ctx, id, status)
}

func (uc *OrderUseCase) GetUserOrders(ctx context.Context, userID primitive.ObjectID) ([]entity.Order, error) {
	return uc.repo.GetUserOrders(ctx, userID)
}

func (uc *OrderUseCase) CancelOrder(ctx context.Context, orderID primitive.ObjectID) error {
	order, err := uc.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	if order.Status == "delivered" {
		return errors.New("cannot cancel delivered order")
	}

	// Process refund if needed
	if order.PaymentStatus == "completed" {
		if err := uc.paymentService.ProcessRefund(ctx, order); err != nil {
			return err
		}
	}

	// Cancel shipping
	if err := uc.shippingService.CancelShipment(ctx, order.ShippingInfo.TrackingID); err != nil {
		return err
	}

	order.Status = "cancelled"
	order.UpdatedAt = time.Now()

	if err := uc.orderRepo.Update(ctx, order); err != nil {
		return err
	}

	return uc.eventPublisher.PublishOrderCancelled(order)
}
