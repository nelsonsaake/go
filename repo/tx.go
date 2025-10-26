package repo

import (
	"gorm.io/gorm"
)

func Tx[T any]() *gorm.DB {
	return newDB().Begin().Model(new(T))
}
