package service

import (
	log "github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/constants"
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"github.com/hthinh24/go-store/services/identity/internal/dto/response"
	"github.com/hthinh24/go-store/services/identity/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	logger         log.Logger
	userRepository identity.UserRepository
	authRepository identity.AuthRepository
}

func NewUserService(logger log.Logger,
	userRepository identity.UserRepository,
	authRepository identity.AuthRepository) identity.UserService {
	return &userService{
		logger:         logger,
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (u *userService) GetUserByID(id int64) (*response.UserResponse, error) {
	u.logger.Info("Get user by ID:", id)

	user, err := u.userRepository.FindUserByID(id)
	if err != nil {
		u.logger.Error("Error fetching user by ID:", err)
		return nil, err
	}

	u.logger.Info("Get user successfully")
	return createUserResponse(user), nil
}

func (u *userService) GetUsers() (*[]response.UserResponse, error) {
	u.logger.Info("Get all users")

	users, err := u.userRepository.FindUsers()
	if err != nil {
		u.logger.Error("Error fetching users:", err)
		return nil, err
	}

	u.logger.Info("Get all users successfully")
	var userResponses []response.UserResponse
	for _, user := range *users {
		userResponses = append(userResponses, *createUserResponse(&user))
	}
	return &userResponses, nil
}

func (u *userService) CreateUser(data *request.CreateUserRequest) (*response.UserResponse, error) {
	u.logger.Info("Creating new user with email:", data.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("Error hashing password:", err)
		return nil, err
	}

	data.Password = string(hashedPassword)
	user := u.createUserEntity(data)
	if err := u.userRepository.CreateUser(user); err != nil {
		u.logger.Error("Error creating user:", err)
		return nil, err
	}

	if err := u.setUserRoleToUser(user); err != nil {
		u.logger.Error("Error setting user role:", err)
		return nil, err
	}

	u.logger.Info("Successfully created user with ID:", user.ID)
	return createUserResponse(user), nil
}

func (u *userService) UpdateUserProfile(id int64, data *request.UpdateUserProfileRequest) (*response.UserResponse, error) {
	user, err := u.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	u.logger.Info("Updating user profile with ID:", id)

	updateUserEntity(user, data)
	if err := u.userRepository.UpdateUserProfile(user); err != nil {
		u.logger.Error("Error updating user profile:", err)
		return nil, err
	}

	u.logger.Info("Successfully updated user profile with ID:", id)
	return createUserResponse(user), nil
}

func (u *userService) UpdateUserPassword(id int64, data *request.UpdateUserPasswordRequest) (*response.UserResponse, error) {
	u.logger.Info("Find user by ID:", id)

	user, err := u.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	u.logger.Info("Updating user password for ID:", id)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.OldPassword)); err != nil {
		u.logger.Error("Old password does not match for user ID:", id)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("Error hashing new password:", err)
		return nil, err
	}

	user.Password = string(hashedPassword)
	if err := u.userRepository.UpdateUserPassword(user); err != nil {
		u.logger.Error("Error updating user password:", err)
		return nil, err
	}
	u.logger.Info("Successfully updated user password for ID:", id)
	return createUserResponse(user), nil
}

func (u *userService) DeleteUser(id int64) error {
	u.logger.Info("Deleting user with ID:", id)

	err := u.userRepository.DeleteUser(id)
	if err != nil {
		u.logger.Error("Error deleting user with ID:", id, "Error:", err)
		return err
	}

	u.logger.Info("Successfully deleted user with ID:", id)
	return nil
}

func (u *userService) createUserEntity(user *request.CreateUserRequest) *entity.User {
	return &entity.User{
		Email:        user.Email,
		Password:     user.Password,
		ProviderID:   user.ProviderID,
		ProviderName: user.ProviderName,
		LastName:     user.LastName,
		FirstName:    user.FirstName,
		Avatar:       user.Avatar,
		Gender:       user.Gender,
		PhoneNumber:  user.PhoneNumber,
		DateOfBirth:  user.DateOfBirth,
		Status:       user.Status,
	}
}

func createUserResponse(user *entity.User) *response.UserResponse {
	return &response.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		LastName:    user.LastName,
		FirstName:   user.FirstName,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Status:      user.Status,
	}
}

func updateUserEntity(user *entity.User, data *request.UpdateUserProfileRequest) {
	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.LastName != nil {
		user.LastName = *data.LastName
	}
	if data.FirstName != nil {
		user.FirstName = *data.FirstName
	}
	if data.Avatar != nil {
		user.Avatar = *data.Avatar
	}
	if data.Gender != nil {
		user.Gender = *data.Gender
	}
	if data.PhoneNumber != nil {
		user.PhoneNumber = *data.PhoneNumber
	}
	if data.DateOfBirth != nil {
		user.DateOfBirth = *data.DateOfBirth
	}
}

func (u *userService) setUserRoleToUser(user *entity.User) error {
	role, err := u.authRepository.FindRoleByName(string(constants.ROLE_USER))
	if err != nil {
		return err
	}

	userHasRole := entity.UserRoles{
		UserID: user.ID,
		RoleID: role.ID,
	}

	if err := u.authRepository.AddRoleToUser(userHasRole); err != nil {
		u.logger.Error("Error assigning role to user:", err)
		return err
	}

	u.logger.Info("Successfully assigned role to user with ID:", user.ID)
	return nil
}
