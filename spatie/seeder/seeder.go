package seeder

import (
	"github.com/nelsonsaake/go/spatie/factory"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	role := factory.FakeRole()
	permission := factory.FakePermission()
	if err := db.Create(role).Error; err != nil {
		return err
	}
	if err := db.Create(permission).Error; err != nil {
		return err
	}
	// Add more seeding logic for other models and relationships
	return nil
}
