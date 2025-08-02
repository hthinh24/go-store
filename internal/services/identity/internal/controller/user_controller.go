package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"net/http"
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
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		u.logger.Info("Fetching user by ID:", id)

		user, err := u.service.GetUserByID(int64(id))
		if err != nil {
			u.logger.Error("Error fetching user by ID:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to fetch user"})
			return
		}

		u.logger.Info("Successfully fetched user by ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User fetched successfully", user))
	}
}

func (u *userController) GetUsers() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		u.logger.Info("Get all users")

		users, err := u.service.GetUsers()
		if err != nil {
			u.logger.Error("Error fetching users:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to fetch users"})
			return
		}

		u.logger.Info("Get all users successfully")
		ctx.JSON(http.StatusOK, users)
	}
}

func (u *userController) CreateUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var userRequest request.CreateUserRequest
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			u.logger.Error("Error binding JSON:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid request body"})
			return
		}

		if err := userRequest.Validate(); err != nil {
			u.logger.Error("Validation fail:")
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.ValidationError, Message: err.Error()})
			return
		}

		u.logger.Info("Creating new user with email:", userRequest.Email)

		user, err := u.service.CreateUser(&userRequest)
		if err != nil {
			u.logger.Error("Error creating user:")
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to create user"})
			return
		}

		u.logger.Info("Successfully created user with ID:", user.ID)
		ctx.JSON(http.StatusCreated, rest.NewAPIResponse(http.StatusCreated, "User created successfully", user))
	}
}

func (u *userController) UpdateUserProfile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		var updateRequest request.UpdateUserProfileRequest
		if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
			u.logger.Error("Error binding JSON:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid request body"})
			return
		}

		// Validate the update request
		if err := updateRequest.Validate(); err != nil {
			u.logger.Error("Validation failed:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.ValidationError, Message: err.Error()})
			return
		}

		u.logger.Info("Updating user profile with ID:", id)

		user, err := u.service.UpdateUserProfile(int64(id), &updateRequest)
		if err != nil {
			u.logger.Error("Error updating user profile:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to update user profile"})
			return
		}

		u.logger.Info("Successfully updated user profile with ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User profile updated successfully", user))
	}
}

func (u *userController) UpdateUserPassword() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		var passwordRequest request.UpdateUserPasswordRequest
		if err := ctx.ShouldBindJSON(&passwordRequest); err != nil {
			u.logger.Error("Error binding JSON:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid request body"})
			return
		}

		if err := passwordRequest.Validate(); err != nil {
			u.logger.Error("Validation failed:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.ValidationError, Message: err.Error()})
			return
		}

		u.logger.Info("Updating user password for ID:", id)

		user, err := u.service.UpdateUserPassword(int64(id), &passwordRequest)
		if err != nil {
			u.logger.Error("Error updating user password:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to update user password"})
			return
		}

		u.logger.Info("Successfully updated user password for ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User password updated successfully", user))
	}
}

func (u *userController) DeleteUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		u.logger.Info("Deleting user with ID:", id)

		if err := u.service.DeleteUser(int64(id)); err != nil {
			u.logger.Error("Error deleting user with ID:", id, "Error:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to delete user"})
			return
		}

		u.logger.Info("Successfully deleted user with ID:", id)
		ctx.JSON(http.StatusNoContent, rest.NewAPIResponse(http.StatusNoContent, "User deleted successfully", nil))
	}
}
