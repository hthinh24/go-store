package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"strconv"
)

type userController struct {
	logger  logger.Logger
	service identity.UserService
}

func NewUserController(logger logger.Logger, service identity.UserService) *userController {
	return &userController{
		logger:  logger,
		service: service,
	}
}

func (u *userController) GetUserByID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		u.logger.Info("Fetching user by ID:", id)

		user, err := u.service.GetUserByID(int64(id))
		if err != nil {
			u.logger.Error("Error fetching user by ID:", err)
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		u.logger.Info("Successfully fetched user by ID:", id)
		ctx.JSON(200, user)
	}
}

func (u *userController) GetUsers() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		u.logger.Info("Get all users")

		users, err := u.service.GetUsers()
		if err != nil {
			u.logger.Error("Error fetching users:", err)
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		u.logger.Info("Get all users successfully")
		ctx.JSON(200, users)
	}
}

func (u *userController) CreateUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var userRequest request.CreateUserRequest
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			u.logger.Error("Error binding JSON:", err)
			ctx.JSON(400, gin.H{"error": "Bad Request"})
			return
		}

		if err := userRequest.Validate(); err != nil {
			u.logger.Error("Validation fail:")
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		u.logger.Info("Creating new user with email:", userRequest.Email)

		user, err := u.service.CreateUser(&userRequest)
		if err != nil {
			u.logger.Error("Error creating user:")
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		u.logger.Info("Successfully created user with ID:", user.ID)
		ctx.JSON(201, user)
	}
}

func (u *userController) UpdateUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		var updateRequest request.UpdateUserProfileRequest
		if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
			u.logger.Error("Error binding JSON:", err)
			ctx.JSON(400, gin.H{"error": "Bad Request"})
			return
		}

		// Validate the update request
		if err := updateRequest.Validate(); err != nil {
			u.logger.Error("Validation failed:", err)
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		u.logger.Info("Updating user profile with ID:", id)

		user, err := u.service.UpdateUser(int64(id), &updateRequest)
		if err != nil {
			u.logger.Error("Error updating user profile:", err)
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		u.logger.Info("Successfully updated user profile with ID:", id)
		ctx.JSON(200, user)
	}
}

func (u *userController) UpdateUserPassword() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		var passwordRequest request.UpdateUserPasswordRequest
		if err := ctx.ShouldBindJSON(&passwordRequest); err != nil {
			u.logger.Error("Error binding JSON:", err)
			ctx.JSON(400, gin.H{"error": "Bad Request"})
			return
		}

		if err := passwordRequest.Validate(); err != nil {
			u.logger.Error("Validation failed:", err)
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		u.logger.Info("Updating user password for ID:", id)

		user, err := u.service.UpdateUserPassword(int64(id), &passwordRequest)
		if err != nil {
			u.logger.Error("Error updating user password:", err)
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		u.logger.Info("Successfully updated user password for ID:", id)
		ctx.JSON(200, user)
	}
}

func (u *userController) DeleteUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		u.logger.Info("Deleting user with ID:", id)

		if err := u.service.DeleteUser(int64(id)); err != nil {
			u.logger.Error("Error deleting user with ID:", id, "Error:", err)
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		u.logger.Info("Successfully deleted user with ID:", id)
		ctx.JSON(204, nil)
	}
}
