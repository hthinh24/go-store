package http

// HTTP holds HTTP server configuration
type HTTP struct {
	ReadTimeout      int    `mapstructure:"read_timeout" env:"HTTP_READ_TIMEOUT"`
	WriteTimeout     int    `mapstructure:"write_timeout" env:"HTTP_WRITE_TIMEOUT"`
	MaxHeaderBytes   int    `mapstructure:"max_header_bytes" env:"HTTP_MAX_HEADER_BYTES"`
	AllowedOrigins   string `mapstructure:"allowed_origins" env:"ALLOWED_ORIGINS"`
	AllowedMethods   string `mapstructure:"allowed_methods" env:"ALLOWED_METHODS"`
	AllowedHeaders   string `mapstructure:"allowed_headers" env:"ALLOWED_HEADERS"`
	ExposedHeaders   string `mapstructure:"exposed_headers" env:"EXPOSED_HEADERS"`
	AllowCredentials bool   `mapstructure:"allow_credentials" env:"ALLOW_CREDENTIALS"`
}

// IsValid checks if HTTP configuration is valid
func (h *HTTP) IsValid() bool {
	return h.ReadTimeout > 0 && h.WriteTimeout > 0
}

// SetDefaults sets default values for HTTP configuration
func (h *HTTP) SetDefaults() {
	if h.ReadTimeout == 0 {
		h.ReadTimeout = 30 // 30 seconds
	}
	if h.WriteTimeout == 0 {
		h.WriteTimeout = 30 // 30 seconds
	}
	if h.MaxHeaderBytes == 0 {
		h.MaxHeaderBytes = 1048576 // 1MB
	}
	if h.AllowedOrigins == "" {
		h.AllowedOrigins = "*"
	}
	if h.AllowedMethods == "" {
		h.AllowedMethods = "GET,POST,PUT,DELETE,OPTIONS"
	}
	if h.AllowedHeaders == "" {
		h.AllowedHeaders = "Origin,Content-Type,Accept,Authorization,X-Requested-With"
	}
}
