package order

import (
	"context"
	"time"
)

type OrderUseCaseImpl struct {
	repo           OrderRepository
	contextTimeout time.Duration
}

func NewOrderUseCase(repo OrderRepository, timeout time.Duration) OrderUseCase {
	return &OrderUseCaseImpl{repo: repo, contextTimeout: timeout}
}

func (tu *OrderUseCaseImpl) ListOrders(c context.Context) ([]Order, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.repo.ListOrders(ctx)
}

func (tu *OrderUseCaseImpl) Create(c context.Context, order *Order) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	// Add business logic here (e.g., validation, pre-processing)
	order.Status = "Pending"
	err := tu.repo.Create(ctx, order)
	if err != nil {
		return err
	}

	// Emit event (OrderCreated)
	return nil
}

func (tu *OrderUseCaseImpl) FetchByCustomerID(c context.Context, customerId string) ([]Order, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.repo.FetchByCustomerID(ctx, customerId)
}
