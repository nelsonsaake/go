package repo

import (
	"gorm.io/gorm"
)

func getTableName[T any](db *gorm.DB) string {
	var stmt = &gorm.Statement{DB: db}

	stmt.Parse(new(T))

	return stmt.Schema.Table
}

func (r *Repo[T]) GetTableName() string {
	return getTableName[T](r.DB)

}

func GetTableName[T any]() string {
	return Do[T]().GetTableName()
}
