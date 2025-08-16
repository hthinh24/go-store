package redis

import (
	"fmt"
	"time"
)

// Redis holds Redis configuration
type Redis struct {
	Host         string `mapstructure:"host" env:"REDIS_HOST"`
	Port         string `mapstructure:"port" env:"REDIS_PORT"`
	Password     string `mapstructure:"password" env:"REDIS_PASSWORD"`
	DB           int    `mapstructure:"db" env:"REDIS_DB"`
	PoolSize     int    `mapstructure:"pool_size" env:"REDIS_POOL_SIZE"`
	MinIdleConns int    `mapstructure:"min_idle_conns" env:"REDIS_MIN_IDLE_CONNS"`
	DialTimeout  int    `mapstructure:"dial_timeout" env:"REDIS_DIAL_TIMEOUT"`
	ReadTimeout  int    `mapstructure:"read_timeout" env:"REDIS_READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"write_timeout" env:"REDIS_WRITE_TIMEOUT"`
	IdleTimeout  int    `mapstructure:"idle_timeout" env:"REDIS_IDLE_TIMEOUT"`

	// Generic key patterns - each service can define their own
	KeyPatterns map[string]string `mapstructure:"key_patterns"`

	// Generic TTL configurations - each service can define their own
	TTL map[string]time.Duration `mapstructure:"ttl"`
}

// GetAddress returns the Redis address
func (r *Redis) GetAddress() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

// IsValid checks if required Redis fields are set
func (r *Redis) IsValid() bool {
	return r.Host != "" && r.Port != ""
}

// SetDefaults sets default values for Redis configuration
func (r *Redis) SetDefaults() {
	if r.PoolSize == 0 {
		r.PoolSize = 10
	}
	if r.MinIdleConns == 0 {
		r.MinIdleConns = 5
	}
	if r.DialTimeout == 0 {
		r.DialTimeout = 5
	}
	if r.ReadTimeout == 0 {
		r.ReadTimeout = 3
	}
	if r.WriteTimeout == 0 {
		r.WriteTimeout = 3
	}
	if r.IdleTimeout == 0 {
		r.IdleTimeout = 300
	}

	// Initialize maps if nil
	if r.KeyPatterns == nil {
		r.KeyPatterns = make(map[string]string)
	}
	if r.TTL == nil {
		r.TTL = make(map[string]time.Duration)
	}
}

// GetKeyPattern returns a key pattern by name, with optional default
func (r *Redis) GetKeyPattern(name string, defaultPattern ...string) string {
	if pattern, exists := r.KeyPatterns[name]; exists {
		return pattern
	}
	if len(defaultPattern) > 0 {
		return defaultPattern[0]
	}
	return ""
}

// GetTTL returns TTL by name, with optional default
func (r *Redis) GetTTL(name string, defaultTTL ...time.Duration) time.Duration {
	if ttl, exists := r.TTL[name]; exists {
		return ttl
	}
	if len(defaultTTL) > 0 {
		return defaultTTL[0]
	}
	return 0
}

// SetKeyPattern sets a key pattern
func (r *Redis) SetKeyPattern(name, pattern string) {
	if r.KeyPatterns == nil {
		r.KeyPatterns = make(map[string]string)
	}
	r.KeyPatterns[name] = pattern
}

// SetTTL sets a TTL value
func (r *Redis) SetTTL(name string, ttl time.Duration) {
	if r.TTL == nil {
		r.TTL = make(map[string]time.Duration)
	}
	r.TTL[name] = ttl
}
