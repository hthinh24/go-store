package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/cart/internal"
	"github.com/hthinh24/go-store/services/cart/internal/dto/request"
	customErr "github.com/hthinh24/go-store/services/cart/internal/errors"
	"net/http"
	"strconv"
)

type CartController struct {
	logger      logger.Logger
	cartService internal.CartService
}

func NewCartController(logger logger.Logger, cartService internal.CartService) *CartController {
	return &CartController{
		logger:      logger,
		cartService: cartService,
	}
}

func (c *CartController) GetCartItemsByUserID() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		userID, err := c.getUserIDFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, rest.NewErrorResponse(rest.UnauthorizedError, "User not authenticated"))
			return
		}

		cartItems, err := c.cartService.FindCartItemsByCartID(userID)
		if err != nil {
			c.handleCartError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "Get cart items by user id successfully", cartItems))
	}
}

func (c *CartController) AddItemToCart() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var item request.AddItemRequest
		if err := ctx.ShouldBindJSON(&item); err != nil {
			ctx.JSON(http.StatusBadRequest, rest.NewErrorResponse(rest.BadRequestError, "Invalid input"))
			return
		}

		userID, err := c.getUserIDFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, rest.NewErrorResponse(rest.UnauthorizedError, err.Error()))
			return
		}

		if err := c.cartService.AddItemToCart(userID, &item); err != nil {
			c.handleCartError(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, rest.NewAPIResponse(http.StatusCreated, "Item added to cart successfully", nil))
	}
}

func (c *CartController) UpdateItemQuantity() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var item request.UpdateItemQuantityRequest
		if err := ctx.ShouldBindJSON(&item); err != nil {
			ctx.JSON(http.StatusBadRequest, rest.NewErrorResponse(rest.BadRequestError, "Invalid input"))
			return
		}

		userID, err := c.getUserIDFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, rest.NewErrorResponse(rest.UnauthorizedError, err.Error()))
			return
		}

		if err := c.cartService.UpdateItemQuantity(userID, item.ItemID, item.Quantity); err != nil {
			c.handleCartError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "Item quantity updated successfully", nil))
	}
}

func (c *CartController) RemoveItemFromCart() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		itemID := ctx.Param("item_id")
		itemIDInt, err := strconv.Atoi(itemID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, rest.NewErrorResponse(rest.BadRequestError, "Invalid item ID format"))
			return
		}

		userID, err := c.getUserIDFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, rest.NewErrorResponse(rest.UnauthorizedError, err.Error()))
			return
		}

		if err := c.cartService.RemoveItemFromCart(userID, int64(itemIDInt)); err != nil {
			c.handleCartError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, rest.NewAPIResponse(http.StatusOK, "Item removed from cart successfully", nil))
	}
}

func (c *CartController) getUserIDFromContext(ctx *gin.Context) (int64, error) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		return 0, errors.New("invalid user ID type")
	}

	return userIDInt64, nil
}

func (c *CartController) handleCartError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, customErr.ErrCartNotFound):
		ctx.JSON(http.StatusNotFound, rest.NewErrorResponse(rest.NotFoundError, err.Error()))
	case errors.Is(err, customErr.ErrCartItemNotFound):
		ctx.JSON(http.StatusNotFound, rest.NewErrorResponse(rest.NotFoundError, err.Error()))
	case errors.Is(err, customErr.ErrCartItemAlreadyExists):
		ctx.JSON(http.StatusConflict, rest.NewErrorResponse(rest.ConflictError, err.Error()))
	default:
		ctx.JSON(http.StatusInternalServerError, rest.NewErrorResponse(rest.InternalServerErrorError, "An unexpected error occurred"))
	}
}
