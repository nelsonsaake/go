package migration

import (
	"github.com/nelsonsaake/go/passport/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Auth{},
	)
}
