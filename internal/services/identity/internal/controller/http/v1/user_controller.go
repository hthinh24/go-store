package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	customErr "github.com/hthinh24/go-store/services/identity/internal/errors"
	"net/http"
	"strconv"
)

type UserController struct {
	logger      logger.Logger
	userService identity.UserService
}

func NewUserController(logger logger.Logger, service identity.UserService) *UserController {
	return &UserController{
		logger:      logger,
		userService: service,
	}
}

func (u *UserController) GetUserByID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		u.logger.Info("Fetching user by ID:", id)

		user, err := u.userService.GetUserByID(int64(id))
		if err != nil {
			if errors.Is(err, customErr.ErrUserNotFound{}) {
				u.logger.Error("User with ID:", id, "not found")
				ctx.JSON(http.StatusNotFound, rest.ErrorResponse{ApiError: rest.NotFoundError, Message: "User not found"})
				return
			}

			u.logger.Error("Error fetching user by ID:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to fetch user"})
			return
		}

		u.logger.Info("Successfully fetched user by ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User fetched successfully", user))
	}
}

func (u *UserController) GetUsers() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		u.logger.Info("Get all users")

		users, err := u.userService.GetUsers()
		if err != nil {
			u.logger.Error("Error fetching users:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to fetch users"})
			return
		}

		u.logger.Info("Get all users successfully")
		ctx.JSON(http.StatusOK, users)
	}
}

func (u *UserController) CreateUser() func(ctx *gin.Context) {
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

		user, err := u.userService.CreateUser(&userRequest)
		if err != nil {
			u.logger.Error("Error creating user:")
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to create user"})
			return
		}

		u.logger.Info("Successfully created user with ID:", user.ID)
		ctx.JSON(http.StatusCreated, rest.NewAPIResponse(http.StatusCreated, "User created successfully", user))
	}
}

func (u *UserController) UpdateUserProfile() func(ctx *gin.Context) {
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

		user, err := u.userService.UpdateUserProfile(int64(id), &updateRequest)
		if err != nil {
			u.logger.Error("Error updating user profile:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to update user profile"})
			return
		}

		u.logger.Info("Successfully updated user profile with ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User profile updated successfully", user))
	}
}

func (u *UserController) UpdateUserPassword() func(ctx *gin.Context) {
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

		user, err := u.userService.UpdateUserPassword(int64(id), &passwordRequest)
		if err != nil {
			u.logger.Error("Error updating user password:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to update user password"})
			return
		}

		u.logger.Info("Successfully updated user password for ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User password updated successfully", user))
	}
}

func (u *UserController) UpdateToMerchantAccount() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		u.logger.Info("Updating user to merchant account with ID:", id)

		if err := u.userService.UpdateToMerchantAccount(int64(id)); err != nil {
			u.logger.Error("Error updating user to merchant account with ID:", id, "Error:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to update user to merchant account"})
			return
		}

		u.logger.Info("Successfully updated user to merchant account with ID:", id)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "User updated to merchant account successfully", nil))
	}
}

func (u *UserController) DeleteUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			u.logger.Error("Invalid user ID:", idStr, "Error:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid user ID format"})
			return
		}

		u.logger.Info("Deleting user with ID:", id)

		if err := u.userService.DeleteUser(int64(id)); err != nil {
			u.logger.Error("Error deleting user with ID:", id, "Error:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Failed to delete user"})
			return
		}

		u.logger.Info("Successfully deleted user with ID:", id)
		ctx.JSON(http.StatusNoContent, rest.NewAPIResponse(http.StatusNoContent, "User deleted successfully", nil))
	}
}
