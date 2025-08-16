package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/identity/internal/errors"
)

// HandleError handles errors and returns appropriate HTTP responses
func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errors.ErrUserNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
	case errors.ErrUserAlreadyExists:
		response := rest.NewErrorResponse(rest.ConflictError, e.Error())
		c.JSON(http.StatusConflict, response)
	case errors.ErrCartCreationFailed:
		response := rest.NewErrorResponse(rest.InternalServerErrorError, e.Error())
		c.JSON(http.StatusInternalServerError, response)
	case errors.ErrInvalidCredentials:
		response := rest.NewErrorResponse(rest.UnauthorizedError, e.Error())
		c.JSON(http.StatusUnauthorized, response)
	case errors.ErrInvalidUserData:
		response := rest.NewErrorResponse(rest.BadRequestError, e.Error())
		c.JSON(http.StatusBadRequest, response)
	case errors.ErrRoleNotFound:
		response := rest.NewErrorResponse(rest.BadRequestError, e.Error())
		c.JSON(http.StatusBadRequest, response)
	case errors.ErrDatabaseTransaction:
		response := rest.NewErrorResponse(rest.InternalServerErrorError, e.Error())
		c.JSON(http.StatusInternalServerError, response)
	case errors.ErrUserNotActive:
		response := rest.NewErrorResponse(rest.ForbiddenError, e.Error())
		c.JSON(http.StatusForbidden, response)
	case errors.ErrPasswordMismatch:
		response := rest.NewErrorResponse(rest.BadRequestError, e.Error())
		c.JSON(http.StatusBadRequest, response)
	default:
		response := rest.NewErrorResponse(rest.InternalServerErrorError, "An unexpected error occurred")
		c.JSON(http.StatusInternalServerError, response)
	}
}
