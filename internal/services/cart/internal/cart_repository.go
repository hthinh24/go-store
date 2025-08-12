package internal

import "github.com/hthinh24/go-store/services/cart/internal/entity"

type CartRepository interface {
	FindCartByUserID(userID int64) (*entity.Cart, error)
	FindCartItemsByCartID(cartID int64) (*[]entity.CartItem, error)

	FindCartItemByID(id int64) (*entity.CartItem, error)

	AddItemToCart(item *entity.CartItem) error
	UpdateItemQuantity(itemID int64, quantity int) error
	RemoveItemFromCart(itemID int64) error
}
