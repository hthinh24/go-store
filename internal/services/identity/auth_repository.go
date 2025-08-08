package identity

import "github.com/hthinh24/go-store/services/identity/internal/entity"

type AuthRepository interface {
	FindRoleByName(name string) (*entity.Role, error)
	FindAllUserRolesByUserID(userID int64) (*[]entity.Role, error)
	FindAllPermissionsByRoleNames(roleNames []string) (*[]entity.Permission, error)
	FindAllPermissionsByRoleIDs(roleIDs []int64) (*[]entity.Permission, error)
	FindPermissionByName(name string) (*entity.Permission, error)
	AddRoleToUser(userRole *entity.UserRoles) error
}
