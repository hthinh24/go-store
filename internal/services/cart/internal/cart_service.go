package internal

import (
	"github.com/hthinh24/go-store/services/cart/internal/dto/request"
	"github.com/hthinh24/go-store/services/cart/internal/dto/response"
)

type CartService interface {
	CreateCart(data *request.CreateCartRequest) (*response.CartResponse, error)
	GetCartItemsByCartID(userID int64) (*[]response.CartItemResponse, error)
	GetCartByUserID(userID int64) (*response.CartResponse, error)
	AddItemToCart(userID int64, item *request.AddItemRequest) error
	UpdateItemQuantity(userID int64, itemID int64, quantity int) error
	RemoveItemFromCart(userID int64, itemID int64) error
}
