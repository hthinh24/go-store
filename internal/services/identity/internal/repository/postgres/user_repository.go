package postgres

import (
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/identity/internal/entity"
	"github.com/hthinh24/go-store/services/identity/internal/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	Logger logger.Logger
	DB     *gorm.DB
}

func NewUserRepository(logger logger.Logger, db *gorm.DB) *userRepository {
	return &userRepository{
		Logger: logger,
		DB:     db,
	}
}

func (u *userRepository) FindUserByID(id int64) (*entity.User, error) {
	u.Logger.Info("Fetching user with ID: %d", id)

	var user entity.User
	if err := u.DB.First(&user, id).Error; err != nil {
		u.Logger.Error("User with id: %d: not found", id, err)
		return nil, errors.ErrUserNotFound{}
	}

	u.Logger.Info("Successfully fetched user with ID: %d", id)
	return &user, nil
}

func (u *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	u.Logger.Info("Fetching user with email: %s", email)

	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		u.Logger.Error("Error fetching user with email %s: %v", email, err)
		return nil, errors.ErrUserNotFound{}
	}

	u.Logger.Info("Successfully fetched user with email: %s", email)
	return &user, nil
}

func (u *userRepository) FindUsers() (*[]entity.User, error) {
	// TODO Pagination and filtering can be added later

	u.Logger.Info("Fetching all users")

	var users []entity.User
	if err := u.DB.Find(&users).Error; err != nil {
		u.Logger.Error("Error fetching users: %v", err)
		return nil, err
	}

	u.Logger.Info("Get users successfully")
	return &users, nil
}

func (u *userRepository) CreateUser(user *entity.User) error {
	u.Logger.Info("Creating data with email: %s", user.Email)

	if err := u.DB.Create(&user).Error; err != nil {
		u.Logger.Error("Error creating data with email: ", user.Email, "error: ", err)
		return err
	}

	u.Logger.Info("User created successfully with ID: %d", user.ID)
	return nil
}

func (u *userRepository) UpdateUserProfile(user *entity.User) error {
	u.Logger.Info("Updating user profile with ID: %d", user.ID)

	err := u.DB.Save(user).Error
	if err != nil {
		u.Logger.Error("Error updating user with ID %d: %v", user.ID, err)
	}

	u.Logger.Info("User with ID %d updated successfully", user.ID)
	return nil
}

func (u *userRepository) UpdateUserPassword(user *entity.User) error {
	u.Logger.Info("Updating password for user with ID: %d", user.ID)

	if err := u.DB.Save(user).Error; err != nil {
		u.Logger.Error("Error updating password for user with ID %d: %v", user.ID, err)
		return err
	}

	u.Logger.Info("Password for user with ID %d updated successfully", user.ID)
	return nil
}

func (u *userRepository) DeleteUser(id int64) error {
	u.Logger.Info("Deleting user with ID: %d", id)

	if err := u.DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		u.Logger.Error("Error deleting user with ID %d: %v", id, err)
		return err
	}

	u.Logger.Info("User with ID %d deleted successfully", id)
	return nil
}
