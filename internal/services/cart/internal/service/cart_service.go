package service

import (
	"context"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/cart/internal"
	"github.com/hthinh24/go-store/services/cart/internal/config"
	"github.com/hthinh24/go-store/services/cart/internal/constants"
	"github.com/hthinh24/go-store/services/cart/internal/controller/http/client"
	"github.com/hthinh24/go-store/services/cart/internal/dto/request"
	"github.com/hthinh24/go-store/services/cart/internal/dto/response"
	"github.com/hthinh24/go-store/services/cart/internal/entity"
	"github.com/hthinh24/go-store/services/cart/internal/errors"
	"time"
)

type cartService struct {
	logger         logger.Logger
	config         config.AppConfig
	cartRepository internal.CartRepository
	productClient  client.ProductClient
}

func NewCartService(logger logger.Logger, cartRepository internal.CartRepository, productClient client.ProductClient) *cartService {
	return &cartService{
		logger:         logger,
		cartRepository: cartRepository,
		productClient:  productClient,
	}
}

func (c *cartService) CreateCart(data *request.CreateCartRequest) (*response.CartResponse, error) {
	c.logger.Info("Creating cart for user ID: ", data.UserID)

	// Check if cart already exists for this user
	existingCart, err := c.cartRepository.FindCartByUserID(data.UserID)
	if err == nil && existingCart != nil {
		c.logger.Info("Cart already exists for user: ", data.UserID, ", cartID: ", existingCart.ID)
		// Return existing cart with empty items
		return c.createCartResponse(existingCart, &[]response.CartItemResponse{}), nil
	}

	// Create cart entity in service layer
	cart := c.createCartEntity(data.UserID)

	// Pass entity to repository for persistence
	if err := c.cartRepository.CreateCart(cart); err != nil {
		c.logger.Error("Failed to create cart for user: ", data.UserID, ", error: ", err)
		return nil, err
	}

	c.logger.Info("Successfully created cart for user: ", data.UserID, ", cartID: ", cart.ID)
	return c.createCartResponse(cart, &[]response.CartItemResponse{}), nil
}

func (c *cartService) GetCartItemsByCartID(cartID int64) (*[]response.CartItemResponse, error) {
	items, err := c.cartRepository.FindCartItemsByCartID(cartID)
	if err != nil {
		c.logger.Error("Failed to find cart items by cart ID: ", cartID, ", error: ", err)
		return nil, err
	}

	// If no items found, return empty cart items response
	if items == nil || len(*items) == 0 {
		c.logger.Info("No items found in cart: ", cartID)
		return &[]response.CartItemResponse{}, nil
	}

	var cartItemResponses []response.CartItemResponse
	for _, item := range *items {
		cartItemResponse := c.createCartItemResponse(&item)
		cartItemResponses = append(cartItemResponses, cartItemResponse)
	}

	return &cartItemResponses, nil
}

func (c *cartService) GetCartByUserID(userID int64) (*response.CartResponse, error) {
	cart, err := c.cartRepository.FindCartByUserID(userID)
	if err != nil {
		c.logger.Error("Failed to get cart by user ID: ", userID, ", error: ", err)
		return nil, err
	}

	cartItemsResponse, err := c.GetCartItemsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}

	return c.createCartResponse(cart, cartItemsResponse), nil
}

func (c *cartService) AddItemToCart(userID int64, newItem *request.AddItemRequest) error {
	c.logger.Info("Adding item to cart for user ID: ", userID)

	// 1. Find the cart by user ID
	cart, err := c.cartRepository.FindCartByUserID(userID)
	if err != nil {
		c.logger.Error("Failed to find cart by user ID: ", userID, ", error: ", err)
		return err
	}

	// 2. Get latest price & status of the product SKU using client
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productSKUResponse, err := c.productClient.GetProductSKUByID(ctx, newItem.ProductSKUID)
	if err != nil {
		c.logger.Error("Failed to get product SKU details: ", newItem.ProductSKUID, ", error: ", err)
		return errors.ErrProductSKUNotFound
	}

	c.logger.Info("Product SKU status: ", productSKUResponse.SKU)
	if productSKUResponse.Status != constants.ProductStatusActive {
		c.logger.Error("Product SKU is not active: ", newItem.ProductSKUID, ", status: ", productSKUResponse.Status)
		return errors.ErrProductSKUNotActive
	}

	// 3. Store the newItem in the cart with latest price
	cartItemEntity := c.createCartItemEntity(cart.ID, productSKUResponse)
	if err := c.cartRepository.AddItemToCart(cartItemEntity); err != nil {
		c.logger.Error("Failed to add item to cart for user: ", userID, ", error: ", err)
		return err
	}

	return nil
}

func (c *cartService) UpdateItemQuantity(userID int64, itemID int64, quantity int) error {
	c.logger.Info("Updating item quantity in cart for user: ", userID, ", itemID: ", itemID, ", quantity: ", quantity)

	// 1. Find the cart by user ID
	if _, err := c.cartRepository.FindCartByUserID(userID); err != nil {
		c.logger.Error("Failed to find cart by user ID: ", userID, ", error: ", err)
		return err
	}

	// 2. Find the cart item by item ID
	if _, err := c.cartRepository.FindCartItemByID(itemID); err != nil {
		c.logger.Error("Failed to find cart item by ID: ", itemID, ", error: ", err)
		return err
	}

	// 3. Update the item quantity in the cart
	if err := c.cartRepository.UpdateItemQuantity(itemID, quantity); err != nil {
		c.logger.Error("Failed to update cart item quantity: ", itemID, ", error: ", err)
		return err
	}

	return nil
}

func (c *cartService) RemoveItemFromCart(userID int64, itemID int64) error {
	c.logger.Info("Removing item from cart for user: ", userID, ", itemID: ", itemID)

	// 1. Find the cart by user ID
	if _, err := c.cartRepository.FindCartByUserID(userID); err != nil {
		c.logger.Error("Failed to find cart by user ID: ", userID, ", error: ", err)
		return err
	}

	// 2. Find the cart item by item ID
	if _, err := c.cartRepository.FindCartItemByID(itemID); err != nil {
		c.logger.Error("Failed to find cart item by ID: ", itemID, ", error: ", err)
		return err
	}

	// 3. Remove the item from the cart
	if err := c.cartRepository.RemoveItemFromCart(itemID); err != nil {
		c.logger.Error("Failed to remove item from cart: ", itemID, ", error: ", err)
		return err
	}

	return nil
}

func (c *cartService) createCartEntity(userID int64) *entity.Cart {
	return &entity.Cart{
		UserID: userID,
		Status: constants.CartStatusActive,
	}
}

func (c *cartService) createCartItemEntity(cartID int64, productSKUResponse *client.ProductSKUDetailResponse) *entity.CartItem {
	return &entity.CartItem{
		CartID:       cartID,
		ProductID:    productSKUResponse.ProductID,
		ProductSKUID: productSKUResponse.ID,
		UnitPrice:    productSKUResponse.Price,
		Status:       productSKUResponse.Status,
	}
}

func (c *cartService) createCartResponse(cart *entity.Cart, cartItemsResponse *[]response.CartItemResponse) *response.CartResponse {
	return &response.CartResponse{
		ID:     cart.ID,
		UserID: cart.UserID,
		Status: cart.Status,
		Items:  cartItemsResponse,
	}
}

func (c *cartService) createCartItemResponse(item *entity.CartItem) response.CartItemResponse {
	return response.CartItemResponse{
		ID:           item.ID,
		ProductID:    item.ProductID,
		ProductSKUID: item.ProductSKUID,
		Quantity:     item.Quantity,
		UnitPrice:    item.UnitPrice,
		TotalPrice:   item.TotalPrice,
		Status:       item.Status,
	}
}
