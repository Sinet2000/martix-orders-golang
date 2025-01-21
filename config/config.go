package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	AppEnv         string
	ServerAddress  string
	Port           string
	ContextTimeout time.Duration
	MongoDB        MongoConfig
}

type MongoConfig struct {
	URI        string
	DBName     string
	User       string
	Password   string
	AuthSource string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	timeout, _ := strconv.Atoi(getEnvOrDefault("CONTEXT_TIMEOUT", "2"))

	return &AppConfig{
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
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
