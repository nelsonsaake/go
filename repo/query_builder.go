package repo

import (
	"strings"
)

// QueryBuilder provides query filtering functionality
type QueryBuilder[T any] struct {
	conditions []string
	values     []any
}

// NewQueryBuilder creates a new QueryBuilder instance
func NewQueryBuilder[T any]() *QueryBuilder[T] {
	return &QueryBuilder[T]{
		conditions: make([]string, 0),
		values:     make([]any, 0),
	}
}

// WhereNot adds a not equal condition
func (w *QueryBuilder[T]) WhereNot(field string, value any) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" != ?")
	w.values = append(w.values, value)
	return w
}

// WhereLike adds a LIKE condition
func (w *QueryBuilder[T]) WhereLike(field string, value string) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" LIKE ?")
	w.values = append(w.values, value)
	return w
}

// WhereIn adds an IN condition
func (w *QueryBuilder[T]) WhereIn(field string, values ...any) *QueryBuilder[T] {
	placeholders := make([]string, len(values))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	w.conditions = append(w.conditions, field+" IN ("+strings.Join(placeholders, ", ")+")")
	w.values = append(w.values, values...)
	return w
}

// IsNull adds an IS NULL condition
func (w *QueryBuilder[T]) IsNull(field string) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" IS NULL")
	return w
}

// IsNotNull adds an IS NOT NULL condition
func (w *QueryBuilder[T]) IsNotNull(field string) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" IS NOT NULL")
	return w
}

// GreaterThan adds a greater than condition
func (w *QueryBuilder[T]) GreaterThan(field string, value any) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" > ?")
	w.values = append(w.values, value)
	return w
}

// LessThan adds a less than condition
func (w *QueryBuilder[T]) LessThan(field string, value any) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" < ?")
	w.values = append(w.values, value)
	return w
}

// GreaterThanOrEqual adds a greater than or equal condition
func (w *QueryBuilder[T]) GreaterThanOrEqual(field string, value any) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" >= ?")
	w.values = append(w.values, value)
	return w
}

// LessThanOrEqual adds a less than or equal condition
func (w *QueryBuilder[T]) LessThanOrEqual(field string, value any) *QueryBuilder[T] {
	w.conditions = append(w.conditions, field+" <= ?")
	w.values = append(w.values, value)
	return w
}

// And combines conditions with AND
func (w *QueryBuilder[T]) And(other *QueryBuilder[T]) *QueryBuilder[T] {
	if len(other.conditions) > 0 {
		w.conditions = append(w.conditions, other.conditions...)
		w.values = append(w.values, other.values...)
	}
	return w
}

// ToSQL returns the SQL WHERE clause and values
func (w *QueryBuilder[T]) ToSQL() (string, []any) {
	if len(w.conditions) == 0 {
		return "", nil
	}
	return strings.Join(w.conditions, " AND "), w.values
}

// HasConditions returns true if there are any conditions
func (w *QueryBuilder[T]) HasConditions() bool {
	return len(w.conditions) > 0
}

func (w *QueryBuilder[T]) Get(arg GetRequest) (*GetResponse[T], error) {

	repo := Do[T]()

	if w.HasConditions() {
		sql, values := w.ToSQL()
		repo.DB = repo.DB.Where(sql, values...)
	}

	return repo.Get(arg)
}
