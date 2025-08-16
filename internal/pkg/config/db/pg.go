package db

import "fmt"

// PG holds PostgreSQL database configuration
type PG struct {
	Host            string `mapstructure:"host" env:"PG_HOST"`
	Port            string `mapstructure:"port" env:"PG_PORT"`
	User            string `mapstructure:"user" env:"PG_USER"`
	Password        string `mapstructure:"password" env:"PG_PASSWORD"`
	Database        string `mapstructure:"database" env:"PG_DATABASE"`
	SSLMode         string `mapstructure:"ssl_mode" env:"PG_SSL_MODE"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" env:"PG_MAX_OPEN_CONNS"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" env:"PG_MAX_IDLE_CONNS"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime" env:"PG_CONN_MAX_LIFETIME"`
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time" env:"PG_CONN_MAX_IDLE_TIME"`
}

// GetDSN returns the PostgreSQL connection string
func (p *PG) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.Database, p.SSLMode)
}

// IsValid checks if all required PostgreSQL fields are set
func (p *PG) IsValid() bool {
	return p.Host != "" && p.Port != "" && p.User != "" && p.Database != ""
}

// SetDefaults sets default values for PostgreSQL configuration
func (p *PG) SetDefaults() {
	if p.Host == "" {
		p.Host = "localhost"
	}
	if p.Port == "" {
		p.Port = "5432"
	}
	if p.SSLMode == "" {
		p.SSLMode = "disable"
	}
	if p.MaxOpenConns == 0 {
		p.MaxOpenConns = 25
	}
	if p.MaxIdleConns == 0 {
		p.MaxIdleConns = 10
	}
	if p.ConnMaxLifetime == 0 {
		p.ConnMaxLifetime = 300 // 5 minutes in seconds
	}
	if p.ConnMaxIdleTime == 0 {
		p.ConnMaxIdleTime = 60 // 1 minute in seconds
	}
}
