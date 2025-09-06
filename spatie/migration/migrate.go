package migration

import (
	"github.com/nelsonsaake/go/spatie/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
		&models.UserPermission{},
		&models.UserPermissionRevoked{},
	)
}
