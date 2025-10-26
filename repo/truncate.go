package repo

func (r *Repo[T]) Truncate() error {
	return r.Exec("TRUNCATE TABLE " + r.tables).Error
}

func Truncate[T any]() error {
	return Do[T]().Truncate()
}
