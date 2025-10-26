package repo

import (
	"gorm.io/gorm"
)

var (
	dbFactory  func() *gorm.DB
	dbInstance *gorm.DB
)

type Config struct {
	DB        *gorm.DB
	dbFactory func() *gorm.DB
}

func Setup(c Config) {
	dbInstance = c.DB
	dbFactory = c.dbFactory
}

func getDB() *gorm.DB {
	return dbInstance
}

func newDB() *gorm.DB {
	if dbFactory == nil {
		panic("DB factory is not set up. Please call repo.Setup first.")
	}
	return dbFactory()
}
