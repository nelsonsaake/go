package repo

import (
	"gorm.io/gorm"
)

// Repo is a generic repository struct for database operations
// on type T, holding the table name and a gorm.DB instance.
type Repo[T any] struct {
	tables string
	*gorm.DB
}

// Do returns a new Repo instance for type T using the default database connection.
func Do[T any]() *Repo[T] {

	db := getDB()
	return New[T](db, getTableName[T](db))
}

// New creates a new Repo instance for type T with the provided database connection
// and table name.
func New[T any](db *gorm.DB, name string) *Repo[T] {

	db = db.Model(new(T))

	return &Repo[T]{DB: db, tables: name}
}
