package main

import (
	"github.com/hthinh24/go-store/internal/pkg/middleware/auth"
	"github.com/hthinh24/go-store/services/product/internal/config"
	"github.com/hthinh24/go-store/services/product/internal/controller"
	repository "github.com/hthinh24/go-store/services/product/internal/infra/repository/postgres"
	"github.com/hthinh24/go-store/services/product/internal/service"
	"github.com/redis/go-redis/v9"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	customLog "github.com/hthinh24/go-store/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configPath := "config.yaml"

	// Load configuration using shared pkg config
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize appLogger
	appLogger := customLog.NewAppLogger(cfg.GetLogLevel())
	appLogger.Info("Starting Product Service...")
	appLogger.Info("Environment: %s", cfg.GetEnvironment())

	// Init Redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisAddress(),
		Password: cfg.GetRedisPassword(), // no password set
		DB:       0,                      // use default DB
		//PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
		DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
	})

	// Initialize database connection
	db, err := initDatabase(cfg)
	if err != nil {
		appLogger.Error("Failed to connect to database: %v", err)
		log.Fatal(err)
	}
	appLogger.Info("Database connected successfully")

	// Initialize repositories
	productRepository := repository.NewProductRepository(
		customLog.WithComponent(cfg.GetLogLevel(), "PRODUCT-REPOSITORY"),
		db)

	// Initialize services
	productService := service.NewProductService(
		customLog.WithComponent(cfg.GetLogLevel(), "PRODUCT-SERVICE"),
		client,
		productRepository)

	// Initialize controllers
	productController := controller.NewProductController(
		customLog.WithComponent(cfg.GetLogLevel(), "PRODUCT-CONTROLLER"),
		productService)

	// Setup router
	router := setupRouter(productController, cfg)

	// Start server
	serverAddr := cfg.GetServerAddress()
	appLogger.Info("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		appLogger.Error("Failed to start server: %v", err)
		log.Fatal(err)
	}
}

func initDatabase(cfg *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.GetDatabaseURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Apply connection pool settings from config
	sqlDB.SetMaxOpenConns(cfg.PG.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.PG.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.PG.ConnMaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.PG.ConnMaxIdleTime) * time.Second)

	return db, nil
}

func setupRouter(productController *controller.ProductController, cfg *config.AppConfig) *gin.Engine {
	router := gin.Default()

	authMiddleware := auth.NewSharedAuthMiddleware(customLog.WithComponent(cfg.GetLogLevel(), "AUTH-MIDDLEWARE"))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		products := v1.Group("/products")
		{
			// Public routes
			products.GET("/:id", productController.GetProductByID())
			products.GET("/:id/detail", productController.GetProductDetailByID())
			products.GET("/skus/:id", productController.GetProductSKUByID())

			// Protected routes
			// TODO - Implement this later
			//products.GET(":userID/products",
			//	authMiddleware.AuthRequired(),
			//	productController.GetProductByUserID())

			products.POST("",
				authMiddleware.AuthRequired(),
				authMiddleware.RequireAnyPermission("product.create"),
				productController.CreateProduct())

			products.POST("/no-sku",
				authMiddleware.AuthRequired(),
				authMiddleware.RequireAnyPermission("product.create"),
				productController.CreateProductWithoutSKU())

			products.DELETE("/:id",
				authMiddleware.AuthRequired(),
				authMiddleware.RequireAnyPermission("product.delete"),
				productController.DeleteProductByID())
		}
	}

	return router
}
