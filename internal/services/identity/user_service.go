package identity

import "github.com/hthinh24/go-store/services/identity/internal/dto/response"

type UserService interface {
	GetUserByID(id int64) (*response.UserResponse, error)
	GetUsers() (*[]response.UserResponse, error)
	CreateUser(user *response.UserResponse) (*response.UserResponse, error)
	DeleteUser(id int64) error
}
