package repo

// Where adds an equality condition
func (w *QueryBuilder[T]) Where(field string, value any) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" = ?")
	w.values = append(w.values, value)
	return w
}

func (r *Repo[T]) Where(field string, value any) *QueryBuilder[T] {
	return NewQueryBuilder[T]().Where(field, value)
}

func Where[T any](field string, value any) *QueryBuilder[T] {
	return NewQueryBuilder[T]().Where(field, value)
}
