package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

// AppConfig holds all application configuration
type AppConfig struct {
	AppEnv         string        `json:"app_env"`
	ServerAddress  string        `json:"server_address"`
	Port           string        `json:"port"`
	ContextTimeout time.Duration `json:"context_timeout"`
	MongoDB        MongoConfig   `json:"mongodb"`
}

// MongoConfig holds MongoDB specific configuration
type MongoConfig struct {
	URI        string `json:"uri"`
	DBName     string `json:"db_name"`
	User       string `json:"user"`
	Password   string `json:"password"`
	AuthSource string `json:"auth_source"`
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		// Only return error if .env file exists but couldn't be loaded
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	timeout, err := strconv.Atoi(getEnvOrDefault("CONTEXT_TIMEOUT", "2"))
	if err != nil {
		return nil, fmt.Errorf("invalid CONTEXT_TIMEOUT value: %w", err)
	}

	config := &AppConfig{
		AppEnv:         getEnvOrDefault("APP_ENV", "development"),
		ServerAddress:  getEnvOrDefault("SERVER_ADDRESS", ":8080"),
		Port:           getEnvOrDefault("PORT", "8080"),
		ContextTimeout: time.Duration(timeout) * time.Second,
		MongoDB: MongoConfig{
			URI:        getEnvOrDefault("MONGO_URI", "mongodb://localhost:27017"),
			DBName:     getEnvOrDefault("MONGO_DB_NAME", "shopGo"),
			User:       getEnvOrDefault("MONGO_USER", "root"),
			Password:   getEnvOrDefault("MONGO_PASSWORD", ""),
			AuthSource: getEnvOrDefault("MONGO_AUTH_SOURCE", "admin"),
		},
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

// validate checks if the configuration is valid
func (c *AppConfig) validate() error {
	if c.AppEnv == "" {
		return fmt.Errorf("APP_ENV cannot be empty")
	}

	if c.ContextTimeout <= 0 {
		return fmt.Errorf("CONTEXT_TIMEOUT must be positive")
	}

	if err := c.MongoDB.validate(); err != nil {
		return fmt.Errorf("mongodb config validation failed: %w", err)
	}

	return nil
}

// validate checks if the MongoDB configuration is valid
func (c *MongoConfig) validate() error {
	if c.URI == "" {
		return fmt.Errorf("MONGO_URI cannot be empty")
	}

	if c.DBName == "" {
		return fmt.Errorf("MONGO_DB_NAME cannot be empty")
	}

	return nil
}

// getEnvOrDefault retrieves an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetMongoURI constructs the complete MongoDB URI with credentials if provided
func (c *MongoConfig) GetMongoURI() string {
	if c.User != "" && c.Password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=%s",
			c.User,
			c.Password,
			c.URI,
			c.DBName,
			c.AuthSource,
		)
	}
	return c.URI
}

// String returns a string representation of AppConfig with sensitive data masked
func (c *AppConfig) String() string {
	return fmt.Sprintf(
		"AppConfig{AppEnv: %s, ServerAddress: %s, Port: %s, ContextTimeout: %s, MongoDB: %s}",
		c.AppEnv,
		c.ServerAddress,
		c.Port,
		c.ContextTimeout,
		c.MongoDB.String(),
	)
}

// String returns a string representation of MongoConfig with sensitive data masked
func (c *MongoConfig) String() string {
	return fmt.Sprintf(
		"MongoConfig{URI: %s, DBName: %s, User: %s, AuthSource: %s}",
		c.URI,
		c.DBName,
		c.User,
		c.AuthSource,
	)
}
