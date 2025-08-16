package config

import (
	"errors"
	"fmt"
	"github.com/hthinh24/go-store/internal/pkg/config/app"
	"github.com/hthinh24/go-store/internal/pkg/config/db"
	"github.com/hthinh24/go-store/internal/pkg/config/http"
	"github.com/hthinh24/go-store/internal/pkg/config/jwt"
	customLog "github.com/hthinh24/go-store/internal/pkg/config/log"
	"github.com/hthinh24/go-store/internal/pkg/config/redis"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Environment string        `mapstructure:"environment" env:"ENV"`
	ServiceName string        `mapstructure:"service_name" env:"SERVICE_NAME"`
	PG          db.PG         `mapstructure:"pg"`
	App         app.App       `mapstructure:"app"`
	HTTP        http.HTTP     `mapstructure:"http"`
	JWT         jwt.JWT       `mapstructure:"jwt"`
	Redis       redis.Redis   `mapstructure:"redis"`
	Log         customLog.Log `mapstructure:"log"`
	Services    Services      `mapstructure:"services"`
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	// Set config file path
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
	}

	// Enable environment variable support
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		log.Println("No config file found, using environment variables and defaults")
	}

	// Unmarshal config
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Set defaults
	config.setDefaults()

	return config, nil
}

// LoadConfigFromEnv loads configuration only from environment variables
func LoadConfigFromEnv(envFile string) (*Config, error) {
	config := &Config{}

	// Load .env file if specified
	if envFile != "" {
		viper.SetConfigFile(envFile)
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Could not load .env file: %v", err)
		}
	}

	// Enable automatic env
	viper.AutomaticEnv()

	// Manually bind environment variables
	bindEnvVars()

	// Unmarshal to struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Set defaults
	config.setDefaults()

	return config, nil
}

func bindEnvVars() {
	// PostgreSQL
	viper.BindEnv("pg.host", "PG_HOST")
	viper.BindEnv("pg.port", "PG_PORT")
	viper.BindEnv("pg.user", "PG_USER")
	viper.BindEnv("pg.password", "PG_PASSWORD")
	viper.BindEnv("pg.database", "PG_DATABASE")
	viper.BindEnv("pg.ssl_mode", "PG_SSL_MODE")
	viper.BindEnv("pg.max_open_conns", "PG_MAX_OPEN_CONNS")
	viper.BindEnv("pg.max_idle_conns", "PG_MAX_IDLE_CONNS")
	viper.BindEnv("pg.conn_max_lifetime", "PG_CONN_MAX_LIFETIME")
	viper.BindEnv("pg.conn_max_idle_time", "PG_CONN_MAX_IDLE_TIME")

	// App
	viper.BindEnv("app.host", "APP_HOST")
	viper.BindEnv("app.port", "APP_PORT")

	// HTTP
	viper.BindEnv("http.read_timeout", "HTTP_READ_TIMEOUT")
	viper.BindEnv("http.write_timeout", "HTTP_WRITE_TIMEOUT")
	viper.BindEnv("http.max_header_bytes", "HTTP_MAX_HEADER_BYTES")
	viper.BindEnv("http.allowed_origins", "ALLOWED_ORIGINS")
	viper.BindEnv("http.allowed_methods", "ALLOWED_METHODS")
	viper.BindEnv("http.allowed_headers", "ALLOWED_HEADERS")
	viper.BindEnv("http.exposed_headers", "EXPOSED_HEADERS")
	viper.BindEnv("http.allow_credentials", "ALLOW_CREDENTIALS")

	// JWT
	viper.BindEnv("jwt.secret", "JWT_SECRET")
	viper.BindEnv("jwt.expiration", "JWT_EXPIRATION")

	// Redis
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")
	viper.BindEnv("redis.pool_size", "REDIS_POOL_SIZE")
	viper.BindEnv("redis.min_idle_conns", "REDIS_MIN_IDLE_CONNS")
	viper.BindEnv("redis.dial_timeout", "REDIS_DIAL_TIMEOUT")
	viper.BindEnv("redis.read_timeout", "REDIS_READ_TIMEOUT")
	viper.BindEnv("redis.write_timeout", "REDIS_WRITE_TIMEOUT")
	viper.BindEnv("redis.idle_timeout", "REDIS_IDLE_TIMEOUT")

	// Log
	viper.BindEnv("log.level", "LOG_LEVEL")

	// Services
	viper.BindEnv("services.user_service_url", "USER_SERVICE_URL")
	viper.BindEnv("services.product_service_url", "PRODUCT_SERVICE_URL")
	viper.BindEnv("services.cart_service_url", "CART_SERVICE_URL")
	viper.BindEnv("services.order_service_url", "ORDER_SERVICE_URL")
	viper.BindEnv("services.gateway_service_url", "GATEWAY_SERVICE_URL")

	// Environment
	viper.BindEnv("environment", "ENV")
	viper.BindEnv("service_name", "SERVICE_NAME")
}

func (c *Config) setDefaults() {
	if c.Environment == "" {
		c.Environment = "development"
	}
	if c.Log.Level == "" {
		c.Log.Level = "info"
	}

	// Set defaults for each component
	c.App.SetDefaults()
	c.HTTP.SetDefaults()
	c.PG.SetDefaults()
	c.Redis.SetDefaults()
}

// Helper methods
func (c *Config) GetDatabaseURL() string {
	return c.PG.GetDSN()
}

func (c *Config) GetServerAddress() string {
	return c.App.GetAddress()
}

func (c *Config) GetRedisAddress() string {
	return c.Redis.GetAddress()
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) IsStaging() bool {
	return c.Environment == "staging"
}

// GetServiceURL returns the URL for a specific service
func (c *Config) GetServiceURL(serviceName string) string {
	switch serviceName {
	case "user":
		return c.Services.IdentityServiceURL
	case "product":
		return c.Services.ProductServiceURL
	case "cart":
		return c.Services.CartServiceURL
	case "order":
		return c.Services.OrderServiceURL
	case "gateway":
		return c.Services.GatewayServiceURL
	default:
		return ""
	}
}

// GetCurrentServicePort returns the current service port based on environment variables
func GetCurrentServicePort() string {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		return port
	}
	return "8080" // default port
}
