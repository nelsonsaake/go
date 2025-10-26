package repo

import "fmt"

// Create inserts a new record of type T into the database using the provided dto.
// Returns the created record or an error if creation fails.
func (r *Repo[T]) Create(dto any) (*T, error) {

	data, err := cast[T](dto)
	if err != nil {
		return nil, fmt.Errorf("error casting dto: %v", err)
	}

	err = r.DB.Create(data).Error

	return data, err
}

// Create is a convenience function that inserts a new record of type T
// using the provided dto and the default repository instance.
func Create[T any](dto any) (*T, error) {
	return Do[T]().Create(dto)
}
