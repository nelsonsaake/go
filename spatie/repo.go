package spatie

import (
	"context"
	"sync"

	"github.com/nelsonsaake/go/spatie/models"
	"github.com/nelsonsaake/go/spatie/repo"
	"gorm.io/gorm"
)

type Repo = repo.Repo

var (
	repoInstance     *Repo
	repoInstanceOnce sync.Once
)

func newRepo(db *gorm.DB, ctx context.Context) *Repo {
	return repo.New(db, ctx)
}

func connect(db *gorm.DB, ctx context.Context) *Repo {
	repoInstanceOnce.Do(func() {
		repoInstance = newRepo(db, ctx)
	})
	return repoInstance
}

// get repo instance or panic if not initialized
func gri() *Repo {
	if repoInstance == nil {
		panic("spatie instance not initialized, call Connect() first")
	}

	return repoInstance
}

func CreateRole(names ...string) ([]*models.Role, error) {
	return gri().CreateRole(names...)
}

func CreatePermission(names ...string) ([]*models.Permission, error) {
	return gri().CreatePermission(names...)
}

func AssignPermissionToRole(roleName string, permNames ...string) error {
	return gri().AssignPermissionToRole(roleName, permNames...)
}

func GiveRoleToUser(userId string, roleNames ...string) error {
	return gri().GiveRoleToUser(userId, roleNames...)
}

func RemoveRoleFromUser(userId string, roleNames ...string) error {
	return gri().RemoveRoleFromUser(userId, roleNames...)
}

func GivePermissionToUser(userId string, permNames ...string) error {
	return gri().GivePermissionToUser(userId, permNames...)
}

func RevokePermissionFromUser(userId string, permNames ...string) error {
	return gri().RevokePermissionFromUser(userId, permNames...)
}

func Is(userId string, roles ...string) (bool, error) {
	return gri().Is(userId, roles...)
}

func IsAny(userId string, roles ...string) (bool, error) {
	return gri().IsAny(userId, roles...)
}

func Can(userId string, perms ...string) (bool, error) {
	return gri().Can(userId, perms...)
}

func CanAny(userId string, perms ...string) (bool, error) {
	return gri().CanAny(userId, perms...)
}

func GetRoles(userId string) ([]string, error) {
	return gri().GetRoles(userId)
}

func GetDetailedPermissions(userId string) (map[string]struct{ Direct, ViaRole, Revoked bool }, error) {
	return gri().GetDetailedPermissions(userId)
}

func GetPermissions(userId string) ([]string, error) {
	return gri().GetPermissions(userId)
}

func Scope(user any) (*repo.Scope, error) {
	return gri().Scope(user)
}

func Scopes(user repo.User) ([]*repo.Scope, error) {
	return gri().Scopes(user)
}

func MappedScope(users ...repo.User) (map[string]*repo.Scope, error) {
	return gri().MappedScope(users...)
}
