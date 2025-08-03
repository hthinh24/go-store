package postgres

import (
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/identity/internal/entity"
	"gorm.io/gorm"
)

type authRepository struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewAuthRepository(logger logger.Logger, db *gorm.DB) *authRepository {
	return &authRepository{
		logger: logger,
		db:     db,
	}
}

func (a *authRepository) FindRoleByName(name string) (*entity.Role, error) {
	a.logger.Info("Finding role by name:", name)

	var role entity.Role
	if err := a.db.Where("name = ?", name).First(&role).Error; err != nil {
		a.logger.Error("Failed to find role by name:", name, "Error:", err)
		return nil, err
	}

	return &role, nil
}

func (a *authRepository) FindUserPermissionsByUserID(userID int64) (*[]entity.Permission, error) {
	a.logger.Info("Finding permissions for user ID:", userID)

	var permissions []entity.Permission
	err := a.db.Table("user_roles").
		Select("permissions.*").
		Joins("JOIN permissions ON user_permissions.permission_id = permissions.id").
		Where("user_permissions.user_id = ?", userID).
		Find(&permissions).Error
	if err != nil {
		a.logger.Error("Failed to find permissions for user ID:", userID, "Error:", err)
		return nil, err
	}

	return &permissions, nil
}

func (a *authRepository) FindAllUserRolesByUserID(userID int64) (*[]entity.Role, error) {
	a.logger.Info("Finding roles for user ID:", userID)

	var roles []entity.Role
	err := a.db.Table("user_roles").
		Select("roles.*").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		a.logger.Error("Failed to find roles for user ID:", userID, "Error:", err)
		return nil, err
	}

	return &roles, nil
}

func (a *authRepository) FindAllPermissionsByRoleNames(roleNames []string) (*[]entity.Permission, error) {
	a.logger.Info("Finding permissions for role names:", roleNames)

	var permissions []entity.Permission
	err := a.db.Table("role_permissions").
		Select("permissions.*").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_name IN ?", roleNames).
		Find(&permissions).Error
	if err != nil {
		a.logger.Error("Failed to find permissions for role names:", roleNames, "Error:", err)
		return nil, err
	}

	return &permissions, nil
}

func (a *authRepository) FindAllPermissionsByRoleIDs(roleIDs []int64) (*[]entity.Permission, error) {
	a.logger.Info("Finding permissions for role IDs:", roleIDs)

	var permissions []entity.Permission
	err := a.db.Table("role_permissions").
		Select("permissions.*").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id IN ?", roleIDs).
		Find(&permissions).Error
	if err != nil {
		a.logger.Error("Failed to find permissions for role IDs:", roleIDs, "Error:", err)
		return nil, err
	}

	return &permissions, nil
}

func (a *authRepository) FindPermissionByName(name string) (*entity.Permission, error) {
	a.logger.Info("Finding permission by name:", name)

	var permission entity.Permission
	if err := a.db.Where("name = ?", name).First(&permission).Error; err != nil {
		a.logger.Error("Failed to find permission by name:", name, "Error:", err)
		return nil, err
	}

	return &permission, nil
}

func (u *authRepository) AddRoleToUser(role entity.UserRoles) error {
	u.logger.Info("Adding role to user with ID: %d", role.UserID)

	if err := u.db.Create(&role).Error; err != nil {
		u.logger.Error("Error adding role to user with ID %d: %v", role.UserID, err)
		return err
	}

	u.logger.Info("Role added to user with ID %d successfully", role.UserID)
	return nil
}
