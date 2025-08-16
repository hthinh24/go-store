package config

// Services holds URLs for inter-service communication
type Services struct {
	IdentityServiceURL     string `mapstructure:"identity_service_url" env:"IDENTITY_SERVICE_URL"`
	ProductServiceURL      string `mapstructure:"product_service_url" env:"PRODUCT_SERVICE_URL"`
	CartServiceURL         string `mapstructure:"cart_service_url" env:"CART_SERVICE_URL"`
	OrderServiceURL        string `mapstructure:"order_service_url" env:"ORDER_SERVICE_URL"`
	GatewayServiceURL      string `mapstructure:"gateway_service_url" env:"GATEWAY_SERVICE_URL"`
	NotificationServiceURL string `mapstructure:"notification_service_url" env:"NOTIFICATION_SERVICE_URL"`
	PaymentServiceURL      string `mapstructure:"payment_service_url" env:"PAYMENT_SERVICE_URL"`
}

// GetServiceURL returns the URL for a specific service
func (s *Services) GetServiceURL(serviceName string) string {
	switch serviceName {
	case "user", "identity":
		return s.IdentityServiceURL
	case "product":
		return s.ProductServiceURL
	case "cart":
		return s.CartServiceURL
	case "order":
		return s.OrderServiceURL
	case "gateway":
		return s.GatewayServiceURL
	case "notification":
		return s.NotificationServiceURL
	case "payment":
		return s.PaymentServiceURL
	default:
		return ""
	}
}

// SetDefaults sets default service URLs based on standard ports
func (s *Services) SetDefaults() {
	if s.IdentityServiceURL == "" {
		s.IdentityServiceURL = "http://localhost:8080"
	}
	if s.ProductServiceURL == "" {
		s.ProductServiceURL = "http://localhost:8081"
	}
	if s.CartServiceURL == "" {
		s.CartServiceURL = "http://localhost:8082"
	}
	if s.OrderServiceURL == "" {
		s.OrderServiceURL = "http://localhost:8083"
	}
	if s.GatewayServiceURL == "" {
		s.GatewayServiceURL = "http://localhost:8000"
	}
}
