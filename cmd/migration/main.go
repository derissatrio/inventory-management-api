package main

import (
	"log"

	"github.com/joho/godotenv"
	"inventory-ticketing-system/infrastructure/config"
	"inventory-ticketing-system/pkg/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database (this will run auto-migration)
	db, err := database.NewDatabase(cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseConnection(db)

	log.Println("Database migration completed successfully")
}