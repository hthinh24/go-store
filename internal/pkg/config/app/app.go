package app

import "fmt"

// App holds basic application configuration
type App struct {
	Host string `mapstructure:"host" env:"APP_HOST"`
	Port string `mapstructure:"port" env:"APP_PORT"`
}

// GetAddress returns the application address
func (a *App) GetAddress() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

// GetPort returns the application port as string
func (a *App) GetPort() string {
	return a.Port
}

// IsValid checks if required app fields are set
func (a *App) IsValid() bool {
	return a.Port != ""
}

// SetDefaults sets default values for App configuration
func (a *App) SetDefaults() {
	if a.Host == "" {
		a.Host = "localhost"
	}
	if a.Port == "" {
		a.Port = "8080"
	}
}
