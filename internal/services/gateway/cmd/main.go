package main

import (
	"github.com/hthinh24/go-store/services/gateway/internal/config"
	"github.com/hthinh24/go-store/services/gateway/internal/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
)

func main() {
	configPath := "config.yaml"

	// Load configuration using shared pkg config
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize logger
	appLogger := logger.WithComponent(cfg.Log.Level, "GATEWAY")

	// Create gateway
	gateway := router.NewGateway(cfg, appLogger)

	// Setup gin router
	r := gin.Default()

	// Setup routes
	gateway.SetupRoutes(r)

	// Start server
	appLogger.Info("Starting Gateway on port", cfg.GetPort())
	if err := r.Run(cfg.GetServerAddress()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
