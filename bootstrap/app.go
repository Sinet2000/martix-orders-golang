package bootstrap

import (
	"context"
	"log"

	"github.com/Sinet2000/Martix-Orders-Go/internal/data"
)

type Application struct {
	AppConfiguration *AppConfiguration
	DbClient         *data.PostgresDB
}

// App initializes the Application instance
func App() *Application {
	app := &Application{}
	app.AppConfiguration = ConfigureApp()

	// Initialize and connect the database client
	dbClient := data.NewPostgresDB()
	err := dbClient.Connect(
		context.Background(),
		app.AppConfiguration.DBHost,
		app.AppConfiguration.DBPort,
		app.AppConfiguration.DBUser,
		app.AppConfiguration.DBPass,
		app.AppConfiguration.DBName,
	)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}
	app.DbClient = dbClient

	return app
}

// CloseDBConnection gracefully closes the database connection
func (app *Application) CloseDBConnection() {
	if err := app.DbClient.Disconnect(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed successfully.")
	}
}
