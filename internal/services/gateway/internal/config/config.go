package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	ApiVersionV1 = "api/v1"
	ApiVersionV2 = "api/v2"
)

type GatewayConfig struct {
	Port               string
	IdentityServiceURL string
	ProductServiceURL  string
	CartServiceURL     string
}

func LoadConfig(fileName string) (*GatewayConfig, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		return nil, fmt.Errorf("error loading %s file: %w", fileName, err)
	}

	config := &GatewayConfig{
		Port:               getEnv("GATEWAY_PORT", "8080"),
		IdentityServiceURL: getEnv("IDENTITY_SERVICE_URL", "http://localhost:8080"),
		ProductServiceURL:  getEnv("PRODUCT_SERVICE_URL", "http://localhost:8081"),
		CartServiceURL:     getEnv("CART_SERVICE_URL", "http://localhost:8082"),
	}

	return config, nil
}

func (c *GatewayConfig) GetServerAddress() string {
	return ":" + c.Port
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
