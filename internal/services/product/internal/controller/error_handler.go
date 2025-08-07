package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	customErr "github.com/hthinh24/go-store/services/product/internal/errors"
)

// ErrorHandler handles different types of errors and returns appropriate HTTP responses
func (pc *ProductController) ErrorHandler(c *gin.Context, err error, defaultMessage string) {
	pc.logger.Error("Error occurred: %v", err)

	// Handle specific error types with appropriate HTTP status codes
	switch e := err.(type) {
	case customErr.ErrProductNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrProductAlreadyExists:
		response := rest.NewErrorResponse(rest.ConflictError, e.Error())
		c.JSON(http.StatusConflict, response)
		return
	case customErr.ErrCategoryNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrBrandNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrUserNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrAttributeNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrOptionNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrOptionValueNotFound:
		response := rest.NewErrorResponse(rest.NotFoundError, e.Error())
		c.JSON(http.StatusNotFound, response)
		return
	case customErr.ErrSKUAlreadyExists:
		response := rest.NewErrorResponse(rest.ConflictError, e.Error())
		c.JSON(http.StatusConflict, response)
		return
	case customErr.ErrInvalidProductData:
		response := rest.NewErrorResponse(rest.BadRequestError, e.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	case customErr.ErrInvalidSKUData:
		response := rest.NewErrorResponse(rest.BadRequestError, e.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	case customErr.ErrDatabaseTransaction:
		response := rest.NewErrorResponse(rest.InternalServerErrorError, e.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	case customErr.ErrDatabaseConnection:
		response := rest.NewErrorResponse(rest.InternalServerErrorError, e.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	default:
		// Generic error for unknown error types
		response := rest.NewErrorResponse(rest.InternalServerErrorError, defaultMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
}
