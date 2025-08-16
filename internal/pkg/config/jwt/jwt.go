package jwt

import "time"

// JWT holds JWT token configuration
type JWT struct {
	Secret     string `mapstructure:"secret" env:"JWT_SECRET"`
	Expiration string `mapstructure:"expiration" env:"JWT_EXPIRATION"`
	Issuer     string `mapstructure:"issuer" env:"JWT_ISSUER"`
	Audience   string `mapstructure:"audience" env:"JWT_AUDIENCE"`
}

// GetExpirationDuration returns the JWT expiration as time.Duration
func (j *JWT) GetExpirationDuration() (time.Duration, error) {
	if j.Expiration == "" {
		return 15 * time.Minute, nil // default 15 minutes
	}
	return time.ParseDuration(j.Expiration)
}

// IsValid checks if required JWT fields are set
func (j *JWT) IsValid() bool {
	return j.Secret != ""
}
