package errors

import "errors"

var (
	ErrCartNotFound          = errors.New("cart not found")
	ErrCartItemNotFound      = errors.New("cart item not found")
	ErrUnauthorized          = errors.New("unauthorized access to cart")
	ErrProductSKUNotFound    = errors.New("product sku not found")
	ErrCartItemAlreadyExists = errors.New("cart item already exists")
	ErrProductSKUNotActive   = errors.New("can`t add item to cart, product sku is not active")
	ErrFetchProductSKUFailed = errors.New("failed to fetch product sku")
)
