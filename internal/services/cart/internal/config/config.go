package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	// Database Configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Server Configuration
	ServerPort string
	ServerHost string

	ProductServiceURL string

	// Log Configuration
	LogLevel string

	// Redis Configuration
	RedisHost     string
	RedisPort     string
	RedisPassword string

	// Environment
	Environment string
}

func LoadConfig(filename string) (*AppConfig, error) {
	// Load .env file in development
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(filename)
		if err != nil {
			// Don't fail if .env file doesn't exist
			fmt.Println("Warning: .env file not found, using system environment variables")
		}
	}

	config := &AppConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "go_store_identity"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerHost: getEnv("SERVER_HOST", "localhost"),

		ProductServiceURL: getEnv("PRODUCT_SERVICE_URL", "http://localhost:8081"),

		LogLevel: getEnv("LOG_LEVEL", "info"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		Environment: getEnv("ENV", "development"),
	}

	return config, nil
}

func (c *AppConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}

func (c *AppConfig) IsProduction() bool {
	return c.Environment == "production"
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
