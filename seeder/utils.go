package seeder

import (
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

// GetRandom fetches a random record from the given table/model.
// modelPtr should be a pointer to a struct representing the model.
func GetRandom[T any](db *gorm.DB) (*T, error) {

	die := func(f string, a ...any) (*T, error) {
		return nil, fmt.Errorf(f, a...)
	}

	var (
		res       = new(T)
		tableName = getTableName[T](db)
		count     int64
	)

	err := db.Table(tableName).Count(&count).Error
	if err != nil {
		return die("failed to count records in %s: %w", tableName, err)
	}

	if count == 0 {
		return die("no records found in %s", tableName)
	}
	offset := rand.Int63n(count)

	err = db.Table(tableName).Offset(int(offset)).Limit(1).Find(res).Error

	return res, err
}

func GetRandomID(db *gorm.DB, tableName string) (string, error) {
	var id string
	err := db.Table(tableName).Select("id").Order("RANDOM()").Limit(1).Scan(&id).Error
	return id, err
}

func getTableName[T any](db *gorm.DB) string {

	var model = new(T)
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	return stmt.Schema.Table
}
