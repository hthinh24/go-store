package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"net/http"
	"strings"
)

type AuthController struct {
	logger      logger.Logger
	authService identity.AuthService
}

func NewAuthController(logger logger.Logger, service identity.AuthService) *AuthController {
	return &AuthController{
		logger:      logger,
		authService: service,
	}
}

func (a *AuthController) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var AuthRequest request.AuthRequest
		if err := ctx.ShouldBindJSON(&AuthRequest); err != nil {
			a.logger.Error("Error binding JSON:", err)
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{ApiError: rest.BadRequestError, Message: "Invalid request body"})
			return
		}

		a.logger.Info("Processing login for user:", AuthRequest.Email)
		authResponse, err := a.authService.Login(AuthRequest)
		if err != nil {
			a.logger.Error("Error during login:", err)
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{ApiError: rest.InternalServerErrorError, Message: "Login failed"})
			return
		}

		a.logger.Info("Login successful for user:", AuthRequest.Email)
		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "Login successful", authResponse))
	}
}

func (a *AuthController) Verify() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			a.logger.Error("Authorization header is missing")
			ctx.JSON(http.StatusUnauthorized, rest.ErrorResponse{ApiError: rest.UnauthorizedError, Message: "Authorization header required"})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		a.logger.Info("Verifying token:", token)
		verifyResponse, err := a.authService.Verify(token)
		if err != nil {
			a.logger.Error("Token verification failed:", err)
			ctx.JSON(http.StatusUnauthorized, rest.ErrorResponse{ApiError: rest.UnauthorizedError, Message: "Invalid token"})
			return
		}

		a.logger.Info("Token verified successfully for user:", verifyResponse.UserID)
		ctx.JSON(http.StatusOK, verifyResponse)
	}
}
