package repo

// Count returns the total number of records of type T
// in the database using the repository instance.
func (r *Repo[T]) Count() (int64, error) {
	var count int64
	err := r.DB.Count(&count).Error
	return count, err
}

// Count is a convenience function that returns the total number of records
// of type T in the database using the default repository instance.
func Count[T any]() (int64, error) {
	return Do[T]().Count()
}
