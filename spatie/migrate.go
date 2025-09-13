package spatie

import (
	"github.com/nelsonsaake/go/spatie/migration"
	"gorm.io/gorm"
)

// Migrate triggers GORM migrations for all Spatie models
func Migrate(db *gorm.DB) error {
	if spatieInstance == nil {
		return nil // or return an error if preferred
	}
	return migration.Migrate(db)
}
