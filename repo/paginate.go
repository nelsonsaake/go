package repo

// Paginate retrieves a slice of items of type T from the database,
// with results paginated by page and pageSize.
func (r *Repo[T]) Paginate(page, pageSize int) ([]T, error) {
	var items []T
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	err := r.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&items).
		Error
	return items, err
}

// Paginate is a convenience function that retrieves paginated items of type T
// using the default repository instance.
func Paginate[T any](page, pageSize int) ([]T, error) {
	return Do[T]().Paginate(page, pageSize)
}
