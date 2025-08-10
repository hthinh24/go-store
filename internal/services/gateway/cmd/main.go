package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/gateway/internal/config"
	"github.com/hthinh24/go-store/services/gateway/internal/router"
)

func main() {
	fileName := ".env"

	// Load configuration
	cfg, _ := config.LoadConfig(fileName)

	// Initialize logger
	appLogger := logger.WithComponent("info", "GATEWAY")

	// Create gateway
	gateway := router.NewGateway(cfg, appLogger)

	// Setup gin router
	r := gin.Default()

	// Setup routes
	gateway.SetupRoutes(r)

	// Start server
	appLogger.Info("Starting Gateway on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
