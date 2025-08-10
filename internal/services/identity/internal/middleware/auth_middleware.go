package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	logger    logger.Logger
	jwtSecret string
}

type JWTClaims struct {
	UserID      int64    `json:"user_id"`
	Email       string   `json:"email"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func NewAuthMiddleware(logger logger.Logger, jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		logger:    logger,
		jwtSecret: jwtSecret,
	}
}

// AuthRequired validates JWT token and sets user info in context
func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "Authorization header required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "Bearer token required",
			})
			c.Abort()
			return
		}

		claims, err := m.ValidateToken(tokenString)
		if err != nil {
			m.logger.Error("Failed to validate JWT token", "error", err)
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "Invalid token",
			})
			c.Abort()
			return
		}

		// If a new token was generated, add it to response headers

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("roles", claims.Roles)
		c.Set("permissions", claims.Permissions)

		c.Next()
	}
}

// RequirePermission checks if user has specific permission from JWT token
func (m *AuthMiddleware) RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, exists := c.Get("permissions")
		if !exists {
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "User not authenticated",
			})
			c.Abort()
			return
		}

		userPermissions := permissions.([]string)
		for _, userPermission := range userPermissions {
			if userPermission == permission {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, rest.ErrorResponse{
			ApiError: rest.ForbiddenError,
			Message:  "Insufficient permissions",
		})
		c.Abort()
	}
}

// RequireRole checks if user has specific role
// NOTE: This middleware should be used AFTER AuthRequired middleware
func (m *AuthMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "User not authenticated",
			})
			c.Abort()
			return
		}

		userRoles := roles.([]string)
		for _, userRole := range userRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, rest.ErrorResponse{
			ApiError: rest.ForbiddenError,
			Message:  "Insufficient role permissions",
		})
		c.Abort()
	}
}

func (m *AuthMiddleware) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Convert claims to our custom JWTClaims type & check validity
	claims, ok := token.Claims.(*JWTClaims)
	if ok {
		return claims, nil
	}

	return nil, jwt.ErrTokenMalformed
}
