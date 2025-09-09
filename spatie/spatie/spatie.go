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

func Connect(db *gorm.DB, ctx context.Context) *Spatie {
	spatieInstanceOnce.Do(func() {
		spatieInstance = New(db, ctx)
	})
	return spatieInstance
}

func Instance() *Spatie {
	if spatieInstance == nil {
		panic("spatie instance not initialized, call Connect() first")
	}

	return spatieInstance
}

// You must call Instance() with your *gorm.DB and context.Background() in your app startup

// Facade methods (examples, add all Spatie methods here)
func CreateRole(names ...string) ([]*models.Role, error) {
	return Instance().CreateRole(names...)
}

func CreatePermission(names ...string) ([]*models.Permission, error) {
	return Instance().CreatePermission(names...)
}

func AssignPermissionToRole(roleName string, permNames ...string) error {
	return Instance().AssignPermissionToRole(roleName, permNames...)
}

func GiveRoleToUser(userId string, roleNames ...string) error {
	return Instance().GiveRoleToUser(userId, roleNames...)
}

func GivePermissionToUser(userId string, permNames ...string) error {
	return Instance().GivePermissionToUser(userId, permNames...)
}

func RevokePermissionFromUser(userId string, permNames ...string) error {
	return Instance().RevokePermissionFromUser(userId, permNames...)
}

func Is(userId string, roles ...string) (bool, error) {
	return Instance().Is(userId, roles...)
}

func IsAny(userId string, roles ...string) (bool, error) {
	return Instance().IsAny(userId, roles...)
}

func Can(userId string, perms ...string) (bool, error) {
	return Instance().Can(userId, perms...)
}

func CanAny(userId string, perms ...string) (bool, error) {
	return Instance().CanAny(userId, perms...)
}

func GetRoles(userId string) ([]string, error) {
	return Instance().GetRoles(userId)
}

func GetDetailedPermissions(userId string) (map[string]struct{ Direct, ViaRole, Revoked bool }, error) {
	return Instance().GetDetailedPermissions(userId)
}

func GetPermissions(userId string) ([]string, error) {
	return Instance().GetPermissions(userId)
}
