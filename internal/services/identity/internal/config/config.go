package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	// Database Configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// JWT Configuration
	JWTSecret           string
	JWTExpiresIn        time.Duration
	JWTRefreshExpiresIn time.Duration

	// Server Configuration
	ServerPort string
	ServerHost string

	// Log Configuration
	LogLevel string

	// Redis Configuration
	RedisHost     string
	RedisPort     string
	RedisPassword string

	// External Services
	CartServiceURL string

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

		JWTSecret:           getEnv("JWT_SECRET", ""),
		JWTExpiresIn:        time.Hour * 24,      // Default to 24 hours
		JWTRefreshExpiresIn: time.Hour * 24 * 30, // Default to 30 days

		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerHost: getEnv("SERVER_HOST", "localhost"),

		LogLevel: getEnv("LOG_LEVEL", "info"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		CartServiceURL: getEnv("CART_SERVICE_URL", ""),

		Environment: getEnv("ENV", "development"),
	}

	// Validate required fields
	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *AppConfig) validate() error {
	if c.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	if len(c.JWTSecret) < 32 {
		return fmt.Errorf("JWT_SECRET must be at least 32 characters long")
	}

	if c.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	return nil
}

func (c *AppConfig) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode)
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
