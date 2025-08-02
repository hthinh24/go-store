package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/identity/internal/controller"
	repository "github.com/hthinh24/go-store/services/identity/internal/repository/postgres"
	"github.com/hthinh24/go-store/services/identity/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	level := "dev"
	appLogger := logger.NewAppLogger(level)

	dsn := "host=localhost user=postgres password=root dbname=postgres port=5001 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		appLogger.Error("Failed to connect to database", "error", err)
		return
	}

	userRepository := repository.NewUserRepository(logger.WithComponent(level, "USER-REPOSITORY"), db)
	userService := service.NewUserService(logger.WithComponent(level, "USER-SERVICE"), userRepository)
	userController := controller.NewUserController(logger.WithComponent(level, "USER-CONTROLLER"), userService)

	r := gin.Default()
	v1 := r.Group("/v1")
	users := v1.Group("/users")

	users.GET("/:id", userController.GetUserByID())
	users.GET("", userController.GetUsers())
	users.POST("", userController.CreateUser())
	users.DELETE("/:id", userController.DeleteUser())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
