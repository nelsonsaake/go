// Package repo provides generic repository utilities for database operations.
package repo

// GetAll retrieves all records of type T from the database using the repository instance.
func (r *Repo[T]) GetAll() ([]T, error) {

	var items = make([]T, 0)
	err := r.DB.Find(&items).Error

	return items, err
}

// GetAll is a convenience function that retrieves all records of type T from the database.
func GetAll[T any]() ([]T, error) {

	return Do[T]().GetAll()
}
