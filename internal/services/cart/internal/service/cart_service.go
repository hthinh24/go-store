package service

import (
	"encoding/json"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/cart/internal"
	"github.com/hthinh24/go-store/services/cart/internal/config"
	"github.com/hthinh24/go-store/services/cart/internal/constants"
	"github.com/hthinh24/go-store/services/cart/internal/dto/http/product"
	"github.com/hthinh24/go-store/services/cart/internal/dto/request"
	"github.com/hthinh24/go-store/services/cart/internal/dto/response"
	"github.com/hthinh24/go-store/services/cart/internal/entity"
	"github.com/hthinh24/go-store/services/cart/internal/errors"
	"net/http"
)

type cartService struct {
	logger         logger.Logger
	config         config.AppConfig
	cartRepository internal.CartRepository
}

func NewCartService(logger logger.Logger, cartRepository internal.CartRepository) *cartService {
	return &cartService{
		logger:         logger,
		cartRepository: cartRepository,
	}
}

func (c *cartService) FindCartItemsByCartID(cartID int64) (*[]response.CartItemResponse, error) {
	items, err := c.cartRepository.FindCartItemsByCartID(cartID)
	if err != nil {
		c.logger.Error("Failed to find cart items by cart ID, ", "cartID: ", cartID, "error", err)
		return nil, err
	}

	// If no items found, return empty cart items response
	if items == nil || len(*items) == 0 {
		c.logger.Info("No items found in cart", "cartID", cartID)
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
		c.logger.Error("Failed to get cart by user ID, ", "userID: ", userID, "error", err)
		return nil, err
	}

	cartItemsResponse, err := c.FindCartItemsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}

	return c.createCartResponse(cart, cartItemsResponse), nil
}

func (c *cartService) AddItemToCart(userID int64, item *request.AddItemRequest) error {
	c.logger.Info("Adding item to cart for user ID", "userID", userID)

	// 1. Find the cart by user ID
	cart, err := c.cartRepository.FindCartByUserID(userID)
	if err != nil {
		c.logger.Error("Failed to find cart by user ID", "userID", userID, "error", err)
		return err
	}

	// 2. Get price & status of the product SKU
	productSKUResponse, err := c.findProductSKUByID(item.ProductSKUID)
	if err != nil {
		return err
	}
	if productSKUResponse.Status != constants.ProductStatusActive {
		c.logger.Error("Product SKU is not active", "productSKUID", item.ProductSKUID, "status", productSKUResponse.Status)
		return errors.ErrProductSKUNotActive
	}

	// 3. Store the item in the cart
	cartItemEntity := c.createCartItemEntity(cart.ID, productSKUResponse, item)
	if err := c.cartRepository.AddItemToCart(cartItemEntity); err != nil {
		c.logger.Error("Failed to add item to cart", "userID", userID, "error", err)
		return err
	}

	return nil
}

func (c *cartService) UpdateItemQuantity(userID int64, itemID int64, quantity int) error {
	c.logger.Info("Updating item quantity in cart", "userID", userID, "itemID", itemID, "quantity", quantity)

	// 1. Find the cart by user ID
	if _, err := c.cartRepository.FindCartByUserID(userID); err != nil {
		c.logger.Error("Failed to find cart by user ID", "userID", userID, "error", err)
		return err
	}

	// 2. Find the cart item by item ID
	if _, err := c.cartRepository.FindCartItemByID(itemID); err != nil {
		c.logger.Error("Failed to find cart item by ID", "itemID", itemID, "error", err)
		return err
	}

	// 3. Update the item quantity in the cart
	if err := c.cartRepository.UpdateItemQuantity(itemID, quantity); err != nil {
		c.logger.Error("Failed to update cart item quantity", "itemID", itemID, "error", err)
		return err
	}

	return nil
}

func (c *cartService) RemoveItemFromCart(userID int64, itemID int64) error {
	c.logger.Info("Removing item from cart", "userID", userID, "itemID", itemID)

	// 1. Find the cart by user ID
	if _, err := c.cartRepository.FindCartByUserID(userID); err != nil {
		c.logger.Error("Failed to find cart by user ID", "userID", userID, "error", err)
		return err
	}

	// 2. Find the cart item by item ID
	if _, err := c.cartRepository.FindCartItemByID(itemID); err != nil {
		c.logger.Error("Failed to find cart item by ID", "itemID", itemID, "error", err)
		return err
	}

	// 3. Remove the item from the cart
	if err := c.cartRepository.RemoveItemFromCart(itemID); err != nil {
		c.logger.Error("Failed to remove item from cart", "itemID", itemID, "error", err)
		return err
	}

	return nil
}

func (c *cartService) createCartItemEntity(cartID int64, productSKUResponse *product.ProductSKUDetailResponse,
	request *request.AddItemRequest) *entity.CartItem {
	return &entity.CartItem{
		CartID:       cartID,
		ProductID:    request.ProductID,
		ProductSKUID: request.ProductSKUID,
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

func (c *cartService) findProductSKUByID(productSKUID int64) (*product.ProductSKUDetailResponse, error) {
	resp, err := http.Get(c.config.ProductServiceURL + "/v1/products/skus/" + string(productSKUID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.ErrProductSKUNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.ErrFetchProductSKUFailed
	}

	var productSKU product.ProductSKUDetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&productSKU); err != nil {
		return nil, err
	}

	return &productSKU, nil
}
