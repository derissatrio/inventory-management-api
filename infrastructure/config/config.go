package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort     string
	DatabaseConfig DatabaseConfig
	JWTSecret      string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		// If .env file doesn't exist, continue with environment variables
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "your-default-secret-key"),
		DatabaseConfig: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "inventory_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	return config, nil
}

func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DatabaseConfig.Host,
		c.DatabaseConfig.Port,
		c.DatabaseConfig.User,
		c.DatabaseConfig.Password,
		c.DatabaseConfig.DBName,
		c.DatabaseConfig.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}