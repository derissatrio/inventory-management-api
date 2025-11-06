package main

import (
	"bufio"
	"log"
	"os"
	"strings"

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

	// Initialize database
	db, err := database.NewDatabase(cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseConnection(db)

	// Get SQL instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Read and execute seed data file
	file, err := os.Open("migrations/000002_seed_data.up.sql")
	if err != nil {
		log.Fatalf("Failed to open seed data file: %v", err)
	}
	defer file.Close()

	// Read file line by line and execute each statement
	scanner := bufio.NewScanner(file)
	var currentStatement strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}

		currentStatement.WriteString(line)
		currentStatement.WriteString(" ")

		// If line ends with semicolon, execute the statement
		if strings.HasSuffix(line, ";") {
			statement := strings.TrimSpace(currentStatement.String())
			if statement != "" {
				log.Printf("Executing: %s", statement[:min(len(statement), 80)]+"...")

				_, err := sqlDB.Exec(statement)
				if err != nil {
					log.Printf("Warning: Failed to execute statement: %v", err)
					log.Printf("Statement was: %s", statement)
				}
			}
			currentStatement.Reset()
		}
	}

	// Execute any remaining statement
	if currentStatement.Len() > 0 {
		statement := strings.TrimSpace(currentStatement.String())
		log.Printf("Executing final statement: %s", statement[:min(len(statement), 80)]+"...")

		_, err := sqlDB.Exec(statement)
		if err != nil {
			log.Printf("Warning: Failed to execute final statement: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading seed data file: %v", err)
	}

	log.Println("✅ Seed data executed successfully!")
	log.Println("📊 Created:")
	log.Println("   - 10 locations")
	log.Println("   - 1 admin user (admin@company.com / admin123)")
	log.Println("   - 8 sample assets")
	log.Println("   - 4 sample tickets")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}