package identity

import (
	"github.com/hthinh24/go-store/services/identity/internal/entity"
)

type UserRepository interface {
	GetUserByID(id int64) (*entity.User, error)
	GetUsers() (*[]entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUserProfile(user *entity.User) error
	UpdateUserPassword(user *entity.User) error
	DeleteUser(id int64) error
}
