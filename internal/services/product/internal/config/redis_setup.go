package config

import (
	pkg "github.com/hthinh24/go-store/internal/pkg/config/redis"
	"github.com/hthinh24/go-store/services/product/internal/constants"
)

// InitProductRedisConfig initializes Redis config with product-specific patterns
func InitProductRedisConfig(baseRedis *pkg.Redis) *pkg.Redis {
	// Get product-specific patterns and TTLs
	keyPatterns, ttlConfig := constants.GetProductRedisConfig()

	// Set the patterns and TTLs in the base config
	for name, pattern := range keyPatterns {
		baseRedis.SetKeyPattern(name, pattern)
	}

	for name, ttl := range ttlConfig {
		baseRedis.SetTTL(name, ttl)
	}

	return baseRedis
}

// Usage example in product service
func SetupProductRedis() *pkg.Redis {
	// Load base Redis config from environment
	baseRedis := &pkg.Redis{
		Host:         "localhost",
		Port:         "6379",
		PoolSize:     15,
		MinIdleConns: 8,
	}

	baseRedis.SetDefaults()

	// Initialize with product-specific patterns
	return InitProductRedisConfig(baseRedis)
}
