package postgres

import (
	"errors"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/cart/internal/entity"
	customErr "github.com/hthinh24/go-store/services/cart/internal/errors"
	"gorm.io/gorm"
)

type CartRepository struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewCartRepository(logger logger.Logger, db *gorm.DB) *CartRepository {
	return &CartRepository{
		logger: logger,
		db:     db,
	}
}

func (c *CartRepository) FindCartByUserID(userID int64) (*entity.Cart, error) {
	c.logger.Info("Getting cart by user ID", "userID", userID)

	var cart entity.Cart
	if err := c.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.logger.Info("No cart found for user ID", "userID", userID)
			return nil, customErr.ErrCartNotFound
		}
		c.logger.Error("Failed to get cart by user ID", "userID", userID, "error", err)
		return nil, err
	}

	return &cart, nil
}

func (c *CartRepository) FindCartItemsByCartID(cartID int64) (*[]entity.CartItem, error) {
	c.logger.Info("Finding cart items by cart ID", "cartID", cartID)

	var items []entity.CartItem
	if err := c.db.Where("cart_id = ?", cartID).Find(&items).Error; err != nil {
		c.logger.Error("Failed to find cart items", "cartID", cartID, "error", err)
		return nil, err
	}
	if len(items) == 0 {
		c.logger.Info("No items found for cart ID", "cartID", cartID)
		return nil, nil
	}

	return &items, nil
}

func (c *CartRepository) FindCartItemByID(id int64) (*entity.CartItem, error) {
	c.logger.Info("Finding cart item by ID", "itemID", id)

	var item entity.CartItem
	if err := c.db.Where("id = ?", id).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.logger.Info("No cart item found for ID", "itemID", id)
			return nil, customErr.ErrCartItemNotFound
		}
		c.logger.Error("Failed to find cart item by ID", "itemID", id, "error", err)
		return nil, err
	}

	return &item, nil
}

func (c *CartRepository) AddItemToCart(item *entity.CartItem) error {
	c.logger.Info("Adding item to cart", "item", item)

	err := c.db.Create(item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.logger.Error("Duplicate cart item", "item", item, "error", err)
			return customErr.ErrCartItemAlreadyExists
		}
		c.logger.Error("Failed to add item to cart", "item", item, "error", err)
		return err
	}

	return nil
}

func (c *CartRepository) UpdateItemQuantity(itemID int64, quantity int) error {
	c.logger.Info("Updating item quantity", "itemID", itemID, "quantity", quantity)

	result := c.db.Table(entity.CartItem{}.TableName()).
		Where("id = ?", itemID).
		Update("quantity", quantity)
	if result.Error != nil {
		c.logger.Error("Failed to update item quantity", "itemID", itemID, "error", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return customErr.ErrCartItemNotFound
	}
	return nil
}

func (c *CartRepository) RemoveItemFromCart(itemID int64) error {
	c.logger.Info("Removing item from cart", "itemID", itemID)

	result := c.db.Table(entity.CartItem{}.TableName()).
		Where("id = ?", itemID).
		Delete(&entity.CartItem{})
	if result.Error != nil {
		c.logger.Error("Failed to remove item from cart", "itemID", itemID, "error", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return customErr.ErrCartItemNotFound
	}
	return nil
}
