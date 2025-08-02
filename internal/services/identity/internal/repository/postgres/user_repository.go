package postgres

import (
	"github.com/hthinh24/go-store/services/identity/internal/entity"
	"gorm.io/gorm"
	"log"
)

type userRepository struct {
	Logger *log.Logger
	DB     *gorm.DB
}

func NewUserRepository(db *gorm.DB, logger *log.Logger) *userRepository {
	return &userRepository{
		Logger: logger,
		DB:     db,
	}
}

func (u *userRepository) GetUserByID(id int64) (*entity.User, error) {
	var user entity.User

	log.Logger{}
	if err := u.DB.First(&user, id).Error; err != nil {
	}
}

func (u *userRepository) GetUsers() (*[]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteUser(id int64) error {
	//TODO implement me
	panic("implement me")
}
