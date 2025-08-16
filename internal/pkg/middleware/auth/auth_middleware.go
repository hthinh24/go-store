package auth

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
)

type SharedAuthMiddleware struct {
	logger logger.Logger
}

func NewSharedAuthMiddleware(logger logger.Logger) *SharedAuthMiddleware {
	return &SharedAuthMiddleware{
		logger: logger,
	}
}

// AuthRequired checks if user is authenticated by validating headers set by gateway
func (m *SharedAuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if user_id header exists (set by gateway after JWT validation)
		userIDHeader := c.GetHeader("X-User-ID")
		if userIDHeader == "" {
			m.logger.Warn("User ID header missing, header: ", "X-User-ID")
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "User not authenticated",
			})
			c.Abort()
			return
		}

		// Parse user_id
		userID, err := strconv.ParseInt(userIDHeader, 10, 64)
		if err != nil {
			m.logger.Error("Invalid user ID header, ", "user_id: ", userIDHeader, ", error: ", err)
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "Invalid user authentication",
			})
			c.Abort()
			return
		}

		// Get other user info from headers
		email := c.GetHeader("X-User-Email")
		rolesHeader := c.GetHeader("X-User-Roles")
		permissionsHeader := c.GetHeader("X-User-Permissions")

		// Parse roles and permissions (comma-separated strings)
		var roles []string
		if rolesHeader != "" {
			roles = strings.Split(rolesHeader, ",")
			// Trim whitespace
			for i, role := range roles {
				roles[i] = strings.TrimSpace(role)
			}
		}

		var permissions []string
		if permissionsHeader != "" {
			permissions = strings.Split(permissionsHeader, ",")
			// Trim whitespace
			for i, perm := range permissions {
				permissions[i] = strings.TrimSpace(perm)
			}
		}

		// Set user info in context for services to use
		c.Set("user_id", userID)
		c.Set("email", email)
		c.Set("roles", roles)
		c.Set("permissions", permissions)

		c.Next()
	}
}

// RequirePermissions checks if user has ALL specified permissions
func (m *SharedAuthMiddleware) RequirePermissions(requiredPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, exists := c.Get("permissions")
		if !exists {
			m.logger.Warn("Permissions not found in context, context_key: ", "permissions")
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "User not authenticated",
			})
			c.Abort()
			return
		}

		userPermissions := permissions.([]string)
		userPermMap := make(map[string]bool)
		for _, perm := range userPermissions {
			userPermMap[perm] = true
		}

		// Check if user has all required permissions
		for _, requiredPerm := range requiredPermissions {
			if !userPermMap[requiredPerm] {
				m.logger.Warn("Access denied - missing permission, required_permission: ",
					requiredPerm, "user_permissions", userPermissions)
				c.JSON(http.StatusForbidden, rest.ErrorResponse{
					ApiError: rest.ForbiddenError,
					Message:  "Insufficient permissions",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// RequireAnyPermission checks if user has ANY of the specified permissions
func (m *SharedAuthMiddleware) RequireAnyPermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPermissions, exists := c.Get("permissions")
		if !exists {
			m.logger.Warn("Permissions not found in context, context_key: ", "permissions")
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "User not authenticated",
			})
			c.Abort()
			return
		}

		userPerms := userPermissions.([]string)
		for _, userPerm := range userPerms {
			for _, requiredPerm := range permissions {
				if userPerm == requiredPerm {
					c.Next()
					return
				}
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
func (m *SharedAuthMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")
		if !exists {
			m.logger.Warn("Roles not found in context, context_key: ", "roles")
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

// RequireAnyRole checks if user has ANY of the specified roles
func (m *SharedAuthMiddleware) RequireAnyRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("roles")
		if !exists {
			m.logger.Warn("Roles not found in context, context_key: ", "roles")
			c.JSON(http.StatusUnauthorized, rest.ErrorResponse{
				ApiError: rest.UnauthorizedError,
				Message:  "User not authenticated",
			})
			c.Abort()
			return
		}

		userRolesList := userRoles.([]string)
		for _, userRole := range userRolesList {
			for _, requiredRole := range roles {
				if userRole == requiredRole {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, rest.ErrorResponse{
			ApiError: rest.ForbiddenError,
			Message:  "Insufficient role permissions",
		})
		c.Abort()
	}
}
