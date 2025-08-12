package main

import (
	"github.com/hthinh24/go-store/internal/pkg/middleware/auth"
	"log"

	"github.com/gin-gonic/gin"
	customLog "github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/cart/internal/config"
	"github.com/hthinh24/go-store/services/cart/internal/controller"
	repository "github.com/hthinh24/go-store/services/cart/internal/infra/postgres"
	"github.com/hthinh24/go-store/services/cart/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fileName := ".env"
	cfg, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	appLogger := customLog.NewAppLogger(cfg.LogLevel)
	appLogger.Info("Starting Cart Service...")
	appLogger.Info("Environment: %s", cfg.Environment)

	db, err := initDatabase(cfg)
	if err != nil {
		appLogger.Error("Failed to connect to database: %v", err)
		log.Fatal(err)
	}
	appLogger.Info("Database connected successfully")

	cartRepository := repository.NewCartRepository(customLog.WithComponent(cfg.LogLevel, "CART-REPOSITORY"), db)
	cartService := service.NewCartService(customLog.WithComponent(cfg.LogLevel, "CART-SERVICE"), cartRepository)
	cartController := controller.NewCartController(customLog.WithComponent(cfg.LogLevel, "CART-CONTROLLER"), cartService)

	router := setupRouter(cartController, cfg)

	serverAddr := ":" + cfg.ServerPort
	appLogger.Info("Cart service starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		appLogger.Error("Failed to start server: %v", err)
		log.Fatal(err)
	}
}

func initDatabase(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=" + cfg.DBSSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupRouter(cartController *controller.CartController, cfg *config.AppConfig) *gin.Engine {
	router := gin.Default()

	authMiddleware := auth.NewSharedAuthMiddleware(customLog.WithComponent(cfg.LogLevel, "AUTH-MIDDLEWARE"))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		cart := v1.Group("/cart")

		// Apply authentication middleware to cart routes
		cart.Use(authMiddleware.AuthRequired())
		{
			cart.GET("", cartController.GetCartItemsByUserID())
			cart.POST("", cartController.AddItemToCart())
			cart.PUT("", cartController.UpdateItemQuantity())
			cart.DELETE("/:item_id", cartController.RemoveItemFromCart())
		}
	}
	return router
}
