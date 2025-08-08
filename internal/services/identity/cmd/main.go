package main

import (
	"errors"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/config"
	"github.com/hthinh24/go-store/services/identity/internal/constants"
	"github.com/hthinh24/go-store/services/identity/internal/controller"
	"github.com/hthinh24/go-store/services/identity/internal/entity"
	customErr "github.com/hthinh24/go-store/services/identity/internal/errors"
	"github.com/hthinh24/go-store/services/identity/internal/middleware"
	repository "github.com/hthinh24/go-store/services/identity/internal/repository/postgres"
	"github.com/hthinh24/go-store/services/identity/internal/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"gorm.io/gorm"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	appLogger := logger.NewAppLogger(cfg.LogLevel)
	appLogger.Info("Starting Identity Service...")
	appLogger.Info("Environment: %s", cfg.Environment)

	// Initialize database connection
	db, err := initDatabase(cfg)
	if err != nil {
		appLogger.Error("Failed to connect to database: %v", err)
		log.Fatal(err)
	}
	appLogger.Info("Database connected successfully")

	// Initialize repositories
	userRepo := repository.NewUserRepository(appLogger, db)
	authRepo := repository.NewAuthRepository(appLogger, db)

	// Initialize services
	authService := service.NewAuthService(appLogger, userRepo, authRepo, cfg)
	userService := service.NewUserService(appLogger, userRepo, authRepo)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(appLogger, authRepo, cfg.JWTSecret)

	// Initialize controllers
	authController := controller.NewAuthController(appLogger, authService)
	userController := controller.NewUserController(appLogger, userService)

	// Setup router
	router := setupRouter(authController, userController, authMiddleware)

	// Initialize user data
	if err := initUserData(userRepo, authRepo); err != nil {
		appLogger.Error("Failed to initialize user data: %v", err)
		log.Fatal(err)
	} else {
		appLogger.Info("User data initialized successfully")
	}

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

	return db, nil
}

func setupRouter(authController *controller.AuthController, userController *controller.UserController, authMiddleware *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		users := api.Group("/users")

		// Public routes
		{
			users.POST("", userController.CreateUser())

			auth.POST("/login", authController.Login())
		}

		auth.Use(authMiddleware.AuthRequired())
		{
			// TODO - Create login, register, logout endpoints
			//auth.POST("/register", authController.Register())
			//auth.POST("/logout", authMiddleware.AuthRequired(), authController.Logout())
		}

		// User routes (protected)
		users.Use(authMiddleware.AuthRequired())
		{
			users.GET(":id", userController.GetUserByID())

			users.PUT("/:id/profile", userController.UpdateUserProfile())
			users.PATCH("/:id/register-merchant", userController.UpdateToMerchantAccount())
			users.PATCH("/:id/password", userController.UpdateUserPassword())

			// Admin only routes
			users.GET("", authMiddleware.RequireRole("admin"), userController.GetUsers())
		}
	}

	return router
}

func initUserData(userRepository identity.UserRepository, authRepository identity.AuthRepository) error {
	user, err := userRepository.FindUserByID(1)
	if user != nil {
		return nil
	}

	// Create admin user if it does not exist
	if errors.Is(err, customErr.ErrUserNotFound{}) {
		user = createAdminUser()
		if err := userRepository.CreateUser(user); err != nil {
			return err
		}
	}

	// Assign admin role to the user
	role, err := authRepository.FindRoleByName(string(constants.RoleAdmin))
	if err != nil {
		return err
	}

	adminRole := entity.UserRoles{
		UserID: user.ID,
		RoleID: role.ID,
	}

	if err := authRepository.AddRoleToUser(&adminRole); err != nil {
		return err
	}

	log.Printf("Admin user initialized: %s", user.Email)
	return nil
}

func createAdminUser() *entity.User {
	password := "admin"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &entity.User{
		Email:        "test@gmail.com",
		Password:     string(hashedPassword),
		ProviderID:   "1",
		ProviderName: "app",
		LastName:     "Admin",
		FirstName:    "Admin",
		Avatar:       "https://example.com/avatar.png",
		Gender:       string(constants.GenderOther),
		PhoneNumber:  "1234567890",
		DateOfBirth:  time.Now(),
		Status:       string(constants.UserStatusActive),
	}
}
