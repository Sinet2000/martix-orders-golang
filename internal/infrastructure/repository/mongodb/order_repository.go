package mongodb

import (
	"context"
	"github.com/Sinet2000/Martix-Orders-Go/internal/domain/entity"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *orderRepository {
	return &orderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *orderRepository) Create(ctx context.Context, order *entity.Order) error {
	logger.Debug("Creating new order",
		zap.Float64("total_amount", order.TotalAmount),
	)

	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		logger.Error("Failed to create order in MongoDB",
			zap.Error(err),
			zap.String("user_id", order.UserID.Hex()),
			zap.Any("order", order),
		)
		return err
	}

	logger.Info("Order created successfully",
		zap.String("order_id", order.ID.Hex()),
		zap.String("user_id", order.UserID.Hex()),
	)
	return nil
}

func (r *orderRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Order, error) {
	var order entity.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		logger.Error("Failed to get order", zap.Error(err))
		return nil, err
	}
	return &order, nil
}
