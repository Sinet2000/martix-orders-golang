package mongodb

import (
	"context"
	"github.com/Sinet2000/Martix-Orders-Go/internal/infrastructure/config"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

type MongoDbContext struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoService(cfg *config.MongoConfig) (*MongoDbContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		AuthSource: cfg.AuthSource,
		Username:   cfg.User,
		Password:   cfg.Password,
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(cfg.URI).
		SetAuth(credential).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error("Failed to ping MongoDB", zap.Error(err))
		_ = client.Disconnect(ctx)
		return nil, err
	}

	logger.Info("Successfully connected to MongoDB")
	return &MongoDbContext{
		Client: client,
		DB:     client.Database(cfg.DBName),
	}, nil
}

func (m *MongoDbContext) Close(ctx context.Context) error {
	if err := m.Client.Disconnect(ctx); err != nil {
		logger.Error("Failed to disconnect from MongoDB", zap.Error(err))
		return err
	}
	logger.Info("Disconnected from MongoDB")
	return nil
}
