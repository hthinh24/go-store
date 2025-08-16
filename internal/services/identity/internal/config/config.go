package config

import (
	"fmt"
	"github.com/hthinh24/go-store/internal/pkg/config"
	"time"
)

type AppConfig struct {
	*config.Config
}

func LoadConfig(configPath string) (*AppConfig, error) {
	// Load shared configuration from pkg
	sharedConfig, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading shared config: %w", err)
	}

	appConfig := &AppConfig{
		Config: sharedConfig,
	}

	return appConfig, nil
}

// Legacy getter methods for backward compatibility
func (c *AppConfig) GetDBHost() string {
	return c.PG.Host
}

func (c *AppConfig) GetDBPort() string {
	return c.PG.Port
}

func (c *AppConfig) GetDBUser() string {
	return c.PG.User
}

func (c *AppConfig) GetDBPassword() string {
	return c.PG.Password
}

func (c *AppConfig) GetDBName() string {
	return c.PG.Database // Updated field name
}

func (c *AppConfig) GetDBSSLMode() string {
	return c.PG.SSLMode // Updated field name
}

func (c *AppConfig) GetJWTSecret() string {
	return c.JWT.Secret
}

func (c *AppConfig) GetJWTExpiresIn() time.Duration {
	duration, _ := time.ParseDuration(c.JWT.Expiration) // Updated field name
	return duration
}

func (c *AppConfig) GetJWTRefreshExpiresIn() time.Duration {
	// Since refresh_expires_in is not in pkg struct, use a default or derive from expiration
	duration, _ := time.ParseDuration("168h") // 7 days default
	return duration
}

func (c *AppConfig) GetServerPort() string {
	return c.App.Port // Updated to use App.Port
}

func (c *AppConfig) GetServerHost() string {
	return c.App.Host // Updated to use App.Host
}

func (c *AppConfig) GetLogLevel() string {
	return c.Log.Level
}

func (c *AppConfig) GetRedisHost() string {
	return c.Redis.Host
}

func (c *AppConfig) GetRedisPort() string {
	return c.Redis.Port
}

func (c *AppConfig) GetRedisPassword() string {
	return c.Redis.Password
}

func (c *AppConfig) GetCartServiceURL() string {
	return c.Services.GetServiceURL("cart")
}

func (c *AppConfig) GetEnvironment() string {
	return c.Environment
}

func (c *AppConfig) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",
		c.PG.Host, c.PG.Port, c.PG.User, c.PG.Password, c.PG.Database, c.PG.SSLMode) // Updated field names
}

func (c *AppConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.GetServerHost(), c.GetServerPort()) // Updated to use GetServerHost and GetServerPort
}
