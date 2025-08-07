package errors

import "fmt"

// Product related errors
type ErrProductNotFound struct{}

func (e ErrProductNotFound) Error() string {
	return "Product not found"
}

type ErrProductAlreadyExists struct {
	Name string
	Slug string
}

func (e ErrProductAlreadyExists) Error() string {
	if e.Slug != "" {
		return fmt.Sprintf("Product with slug '%s' already exists", e.Slug)
	}
	return fmt.Sprintf("Product with name '%s' already exists", e.Name)
}

type ErrInvalidProductData struct {
	Field   string
	Message string
}

func (e ErrInvalidProductData) Error() string {
	return fmt.Sprintf("Invalid product data - %s: %s", e.Field, e.Message)
}

// Category related errors
type ErrCategoryNotFound struct {
	ID int64
}

func (e ErrCategoryNotFound) Error() string {
	return fmt.Sprintf("Category with ID %d not found", e.ID)
}

// Brand related errors
type ErrBrandNotFound struct {
	ID int64
}

func (e ErrBrandNotFound) Error() string {
	return fmt.Sprintf("Brand with ID %d not found", e.ID)
}

// Attribute related errors
type ErrAttributeNotFound struct {
	ID int64
}

func (e ErrAttributeNotFound) Error() string {
	return fmt.Sprintf("Product attribute with ID %d not found", e.ID)
}

// Option related errors
type ErrOptionNotFound struct {
	ID int64
}

func (e ErrOptionNotFound) Error() string {
	return fmt.Sprintf("Product option with ID %d not found", e.ID)
}

type ErrOptionValueNotFound struct {
	Value    string
	OptionID int64
}

func (e ErrOptionValueNotFound) Error() string {
	return fmt.Sprintf("Option value '%s' not found for option ID %d", e.Value, e.OptionID)
}

// SKU related errors
type ErrSKUAlreadyExists struct {
	SKU string
}

func (e ErrSKUAlreadyExists) Error() string {
	return fmt.Sprintf("SKU '%s' already exists", e.SKU)
}

type ErrInvalidSKUData struct {
	SKU     string
	Message string
}

func (e ErrInvalidSKUData) Error() string {
	return fmt.Sprintf("Invalid SKU data for '%s': %s", e.SKU, e.Message)
}

// User related errors
type ErrUserNotFound struct {
	ID int64
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("User with ID %d not found", e.ID)
}

// Database related errors
type ErrDatabaseConnection struct{}

func (e ErrDatabaseConnection) Error() string {
	return "Database connection failed"
}

type ErrDatabaseTransaction struct {
	Operation string
}

func (e ErrDatabaseTransaction) Error() string {
	return fmt.Sprintf("Database transaction failed during %s", e.Operation)
}
