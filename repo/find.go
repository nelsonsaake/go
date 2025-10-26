package repo

// Find searches for a record of type T in the database by its id.
// Returns the found record or an error.
func (r *Repo[T]) Find(id string) (*T, error) {

	var res = new(T)

	err := r.DB.Where("id = ?", id).First(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Find is a convenience function that searches for a record of type T
// by id using the default repository instance.
func Find[T any](id string) (*T, error) {

	return Do[T]().Find(id)
}
