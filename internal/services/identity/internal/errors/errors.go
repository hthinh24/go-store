package errors

import "fmt"

// User related errors
type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "User not found"
}

type ErrUserAlreadyExists struct {
	Email string
}

func (e ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("User with email '%s' already exists", e.Email)
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return "Invalid email or password"
}

type ErrInvalidUserData struct {
	Field   string
	Message string
}

func (e ErrInvalidUserData) Error() string {
	return fmt.Sprintf("Invalid user data - %s: %s", e.Field, e.Message)
}

// Role related errors
type ErrRoleNotFound struct {
	Name string
}

func (e ErrRoleNotFound) Error() string {
	return fmt.Sprintf("Role '%s' not found", e.Name)
}

// Database related errors
type ErrDatabaseTransaction struct {
	Operation string
}

func (e ErrDatabaseTransaction) Error() string {
	return fmt.Sprintf("Database transaction failed during %s", e.Operation)
}

type ErrUserNotActive struct{}

func (e ErrUserNotActive) Error() string {
	return "User is not active"
}

type ErrPasswordMismatch struct{}

func (p ErrPasswordMismatch) Error() string {
	return "Password & confirm password not match"
}
