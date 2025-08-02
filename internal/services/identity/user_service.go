package identity

import (
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"github.com/hthinh24/go-store/services/identity/internal/dto/response"
)

type UserService interface {
	GetUserByID(id int64) (*response.UserResponse, error)
	GetUsers() (*[]response.UserResponse, error)
	CreateUser(data *request.CreateUserRequest) (*response.UserResponse, error)
	UpdateUser(id int64, data *request.UpdateUserProfileRequest) (*response.UserResponse, error)
	UpdateUserPassword(id int64, data *request.UpdateUserPasswordRequest) (*response.UserResponse, error)
	DeleteUser(id int64) error
}
