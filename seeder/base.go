package seeder

import (
	"fmt"

	"gorm.io/gorm"
)

type seeder[T any] struct {
	Table   string
	DB      *gorm.DB
	Factory Factory[T]
}

func (s seeder[T]) getData(args ...int) []T {

	var count = 1
	if len(args) > 0 {
		count = args[0]
	}

	var data = []T{}
	for range count {
		data = append(data, s.Factory())
	}

	return data
}

func (s seeder[T]) Seed(args ...int) error {

	data := s.getData(args...)

	if len(data) == 0 {
		return fmt.Errorf("no data to seed")
	}

	if s.DB == nil {
		return fmt.Errorf("database connection is nil")
	}

	err := s.DB.Table(s.Table).Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func New[T any](db *gorm.DB, factory Factory[T]) seeder[T] {
	return seeder[T]{
		Table:   getTableName[T](db),
		DB:      db,
		Factory: factory,
	}
}
