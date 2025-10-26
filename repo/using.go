package repo

import "gorm.io/gorm"

func Using[T any](db *gorm.DB) *Repo[T] {
	return New[T](db, getTypeName(new(T)))
}
