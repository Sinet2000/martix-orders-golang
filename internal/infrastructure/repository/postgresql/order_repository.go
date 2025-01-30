package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/entity"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type orderRepository struct {
	db *pgxpool.Pool
	// logger *zap.Logger
}

func NewOrderRepository(db *pgxpool.Pool) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("failed to begin transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback(ctx)

	// Insert order
	query := `
        INSERT INTO orders (user_id, total_amount, status, shipping_addr, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	err = tx.QueryRow(ctx, query,
		order.UserID,
		order.TotalAmount,
		order.Status,
		order.ShippingAddr,
		time.Now(),
		time.Now(),
	).Scan(&order.ID)

	if err != nil {
		logger.Error("failed to insert order", zap.Error(err))
		return err
	}

	// Insert order items
	for i := range order.Items {
		query := `
            INSERT INTO order_items (order_id, product_id, quantity, unit_price, total_price)
            VALUES ($1, $2, $3, $4, $5)
            RETURNING id`

		err = tx.QueryRow(ctx, query,
			order.ID,
			order.Items[i].ProductID,
			order.Items[i].Quantity,
			order.Items[i].UnitPrice,
			order.Items[i].TotalPrice,
		).Scan(&order.Items[i].ID)

		if err != nil {
			logger.Error("failed to insert order item",
				zap.Error(err),
				zap.Int64("orderId", order.ID),
				zap.Int64("productId", order.Items[i].ProductID))
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *orderRepository) GetOrderByID(ctx context.Context, orderID int64) (*entity.Order, error) {
	order := &entity.Order{}

	// Get order details
	query := `
        SELECT id, user_id, total_amount, status, shipping_addr, created_at, updated_at
        FROM orders
        WHERE id = $1`

	err := r.db.QueryRow(ctx, query, orderID).Scan(
		&order.ID,
		&order.UserID,
		&order.TotalAmount,
		&order.Status,
		&order.ShippingAddr,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("order not found: %d", orderID)
		}
		logger.Error("failed to get order", zap.Error(err), zap.Int64("orderId", orderID))
		return nil, err
	}

	// Get order items
	query = `
        SELECT id, product_id, quantity, unit_price, total_price
        FROM order_items
        WHERE order_id = $1`

	rows, err := r.db.Query(ctx, query, orderID)
	if err != nil {
		logger.Error("failed to get order items", zap.Error(err), zap.Int64("orderId", orderID))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.OrderItem
		err := rows.Scan(
			&item.ID,
			&item.ProductID,
			&item.Quantity,
			&item.UnitPrice,
			&item.TotalPrice,
		)
		if err != nil {
			logger.Error("failed to scan order item", zap.Error(err))
			return nil, err
		}
		item.OrderID = orderID
		order.Items = append(order.Items, item)
	}

	return order, nil
}
