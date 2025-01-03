package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type PostgresDB struct {
	client *sql.DB
}

// NewPostgresDB initializes a new PostgresDB instance
func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

// Connect initializes the connection to the PostgreSQL database
func (p *PostgresDB) Connect(ctx context.Context, dbHost string, dbPort int, dbUser, dbPass, dbName string) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set connection pooling configurations
	// db.SetMaxOpenConns(maxConns)
	// db.SetMaxIdleConns(maxConns / 2)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Ensure the database connection is valid
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to the PostgreSQL database successfully.")
	p.client = db
	return nil
}

// Disconnect closes the connection to the PostgreSQL database
func (p *PostgresDB) Disconnect() error {
	if p.client != nil {
		if err := p.client.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
		log.Println("PostgreSQL database connection closed successfully.")
	}
	return nil
}

// Client returns the raw *sql.DB instance
func (p *PostgresDB) Client() *sql.DB {
	return p.client
}
