package config

import (
	"fmt"
	"github.com/hthinh24/go-store/internal/pkg/config"
)

const (
	ApiVersionV1 = "api/v1"
	ApiVersionV2 = "api/v2"
)

type GatewayConfig struct {
	*config.Config
}

func LoadConfig(configPath string) (*GatewayConfig, error) {
	// Load shared configuration from pkg
	sharedConfig, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading shared config: %w", err)
	}

	gatewayConfig := &GatewayConfig{
		Config: sharedConfig,
	}

	return gatewayConfig, nil
}

func (c *GatewayConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.App.Host, c.App.Port)
}

// Legacy methods for backward compatibility
func (c *GatewayConfig) GetPort() string {
	return c.App.Port // Updated to use App.Port
}

func (c *GatewayConfig) GetIdentityServiceURL() string {
	return c.Services.GetServiceURL("identity")
}

func (c *GatewayConfig) GetProductServiceURL() string {
	return c.Services.GetServiceURL("product")
}

func (c *GatewayConfig) GetCartServiceURL() string {
	return c.Services.GetServiceURL("cart")
}
