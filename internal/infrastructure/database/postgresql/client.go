package postgresql

import (
	"context"
	"fmt"
	"github.com/Sinet2000/Martix-Orders-Go/internal/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type PostgresqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	HostName string `json:"host_name"`
	Port     string `json:"port"`
}

type PostgresContext struct {
	Pool *pgxpool.Pool
}

func NewPostgresService(cfg *PostgresqlConfig) (*PostgresContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(cfg.GetPostgresURI())
	if err != nil {
		logger.Error("Failed to parse PostgreSQL config", zap.Error(err))
		return nil, err
	}

	// Set additional pool configuration if needed
	config.MaxConns = 10
	config.MinConns = 1
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL", zap.Error(err))
		return nil, err
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		logger.Error("Failed to ping PostgreSQL", zap.Error(err))
		pool.Close()
		return nil, err
	}

	logger.Info("Successfully connected to PostgreSQL")
	return &PostgresContext{
		Pool: pool,
	}, nil
}

func (p *PostgresContext) Close() {
	if p.Pool != nil {
		p.Pool.Close()
		logger.Info("Disconnected from PostgreSQL")
	}
}

func (c *PostgresqlConfig) validate() error {
	if c.Username == "" {
		return fmt.Errorf("POSTGRES_USER cannot be empty")
	}

	if c.Password == "" {
		return fmt.Errorf("POSTGRES_PASSWORD cannot be empty")
	}

	return nil
}

func (c *PostgresqlConfig) GetPostgresURI() string {
	if c.Username != "" && c.Password != "" {
		return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
			c.Username,
			c.Password,
			c.HostName,
			c.Port,
			c.DbName,
		)
	}

	return ""
}
