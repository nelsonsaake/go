package repo

// Exists searches for a record of type T in the database where the specified key
// matches the provided value. Returns a pointer to the found record or an error.
//
// Example:
//   user, err := repo.Exists[User]("email", "user@example.com")
//
// Parameters:
//   key   - the column name to search by
//   value - the value to match
//
// Returns:
//   *T    - pointer to the found record (or nil if not found)
//   error - error if query fails
// Exists checks if any record of type T exists in the database where the specified key matches the provided value.
// Returns true if at least one record exists, false otherwise, and an error if the query fails.
func (r *Repo[T]) Exists(key string, value string) (bool, error) {
	var count int64
	tx := r.DB.Model(new(T)).Where(key+" = ?", value).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}

// Exists is a convenience function that searches for a record of type T
// by key and value using the default repository instance.
//
// Example:
//   user, err := repo.Exists[User]("email", "user@example.com")
//
// Parameters:
//   key   - the column name to search by
//   value - the value to match
//
// Returns:
//   *T    - pointer to the found record (or nil if not found)
//   error - error if query fails
// Exists is a convenience function that checks if any record of type T exists by key and value using the default repository instance.
func Exists[T any](key string, value string) (bool, error) {
	return Do[T]().Exists(key, value)
}
