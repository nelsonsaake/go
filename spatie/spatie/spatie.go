package spatie

import (
	"context"
	"sync"

	"github.com/nelsonsaake/go/spatie/engine"
	"github.com/nelsonsaake/go/spatie/models"
	"gorm.io/gorm"
)

type Spatie = engine.Spatie

func New(db *gorm.DB, ctx context.Context) *Spatie {
	return engine.New(db, ctx)
}

// Singleton management

var (
	spatieInstance     *Spatie
	spatieInstanceOnce sync.Once
)

func Instance(db *gorm.DB, ctx context.Context) *Spatie {
	spatieInstanceOnce.Do(func() {
		spatieInstance = New(db, ctx)
	})
	return spatieInstance
}

// You must call Instance() with your *gorm.DB and context.Background() in your app startup

// Facade methods (examples, add all Spatie methods here)
func CreateRole(names ...string) ([]*models.Role, error) {
	return spatieInstance.CreateRole(names...)
}

func CreatePermission(names ...string) ([]*models.Permission, error) {
	return spatieInstance.CreatePermission(names...)
}

func AssignPermissionToRole(roleName string, permNames ...string) error {
	return spatieInstance.AssignPermissionToRole(roleName, permNames...)
}

func GiveRoleToUser(userId string, roleNames ...string) error {
	return spatieInstance.GiveRoleToUser(userId, roleNames...)
}

func GivePermissionToUser(userId string, permNames ...string) error {
	return spatieInstance.GivePermissionToUser(userId, permNames...)
}

func RevokePermissionFromUser(userId string, permNames ...string) error {
	return spatieInstance.RevokePermissionFromUser(userId, permNames...)
}

func Is(userId string, roles ...string) (bool, error) {
	return spatieInstance.Is(userId, roles...)
}

func IsAny(userId string, roles ...string) (bool, error) {
	return spatieInstance.IsAny(userId, roles...)
}

func Can(userId string, perms ...string) (bool, error) {
	return spatieInstance.Can(userId, perms...)
}

func CanAny(userId string, perms ...string) (bool, error) {
	return spatieInstance.CanAny(userId, perms...)
}

func GetRolesForUser(userId string) ([]string, error) {
	return spatieInstance.GetRolesForUser(userId)
}

func GetDetailedPermissionsForUser(userId string) (map[string]struct{ Direct, ViaRole, Revoked bool }, error) {
	return spatieInstance.GetDetailedPermissionsForUser(userId)
}

func GetPermissionsForUser(userId string) ([]string, error) {
	return spatieInstance.GetPermissionsForUser(userId)
}
