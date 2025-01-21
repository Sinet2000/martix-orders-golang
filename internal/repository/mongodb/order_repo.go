package mongodb

import (
	"context"
	"errors"
	"github.com/Sinet2000/Martix-Orders-Go/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type OrderRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewOrderRepository(db *mongo.Database, logger *zap.Logger) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
		logger:     logger,
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *entity.Order) error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		r.logger.Error("Failed to create order", zap.Error(err))
		return err
	}

	order.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *OrderRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Order, error) {
	var order entity.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		r.logger.Error("Failed to get order", zap.Error(err))
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	update := bson.M{
		"$set": bson.M{
			"status":    status,
			"updatedAt": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		r.logger.Error("Failed to update order status", zap.Error(err))
		return err
	}
	return nil
}

func (r *OrderRepository) GetUserOrders(ctx context.Context, userID primitive.ObjectID) ([]entity.Order, error) {
	var orders []entity.Order
	cursor, err := r.collection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		r.logger.Error("Failed to get user orders", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &orders); err != nil {
		r.logger.Error("Failed to decode orders", zap.Error(err))
		return nil, err
	}
	return orders, nil
}
