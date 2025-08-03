package identity

import (
	"github.com/hthinh24/go-store/services/identity/internal/entity"
)

type UserRepository interface {
	FindUserByID(id int64) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindUsers() (*[]entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUserProfile(user *entity.User) error
	UpdateUserPassword(user *entity.User) error
	DeleteUser(id int64) error
}
