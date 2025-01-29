package repository

import (
	"context"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context, userID primitive.ObjectID) ([]entity.Order, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status entity.OrderStatus) error
}
