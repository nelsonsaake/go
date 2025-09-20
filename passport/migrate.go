package passport

import (
	"github.com/nelsonsaake/go/passport/migration"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return migration.Migrate(db)
}
