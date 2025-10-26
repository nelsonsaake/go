package repo

// FindBy searches for a record of type T in the database where the given key
// matches the provided value. Returns the found record or an error.
func (r *Repo[T]) FindBy(key string, value string) (*T, error) {

	var res = new(T)

	err := r.DB.Where(key+" = ?", value).Find(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// FindBy is a convenience function that searches for a record of type T
// by key and value using the default repository instance.
func FindBy[T any](key string, value string) (*T, error) {

	return Do[T]().FindBy(key, value)
}
