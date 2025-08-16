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
	return c.PG.Database
}

func (c *AppConfig) GetDBSSLMode() string {
	return c.PG.SSLMode
}

func (c *AppConfig) GetJWTSecret() string {
	return c.JWT.Secret
}

func (c *AppConfig) GetJWTExpire() string {
	return c.JWT.Expiration
}

func (c *AppConfig) GetJWTExpiresIn() time.Duration {
	duration, _ := time.ParseDuration(c.JWT.Expiration)
	return duration
}

func (c *AppConfig) GetServerPort() string {
	return c.App.Port
}

func (c *AppConfig) GetServerHost() string {
	return c.App.Host
}

func (c *AppConfig) GetUserServiceURL() string {
	return c.Services.GetServiceURL("identity")
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

func (c *AppConfig) GetEnvironment() string {
	return c.Environment
}

func (c *AppConfig) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		c.PG.Host, c.PG.User, c.PG.Password, c.PG.Database, c.PG.Port, c.PG.SSLMode)
}

func (c *AppConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.App.Host, c.App.Port)
}

func (c *AppConfig) IsProduction() bool {
	return c.Environment == "production"
}
