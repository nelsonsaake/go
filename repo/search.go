package repo

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/nelsonsaake/go/strs"
)

var (
	searchableFieldsCache = make(map[reflect.Type][]string)
	searchableFieldsMutex sync.RWMutex
)

// getSearchableFields takes any value, marshals it to JSON,
// and returns the keys whose values are strings.
func getSearchableFields[T any]() ([]string, error) {
	typ := reflect.TypeOf((*T)(nil)).Elem()

	// Check cache first
	searchableFieldsMutex.RLock()
	if fields, exists := searchableFieldsCache[typ]; exists {
		searchableFieldsMutex.RUnlock()
		return fields, nil
	}
	searchableFieldsMutex.RUnlock()

	// Compute if not in cache
	m, err := toMap(new(T))
	if err != nil {
		return nil, err
	}

	fields := make([]string, 0)
	for name, v := range m {
		if _, ok := v.(string); ok {
			// db field names are in snake_case
			fields = append(fields, strs.ToSnakeCase(name))
		}
	}

	// Store in cache
	searchableFieldsMutex.Lock()
	searchableFieldsCache[typ] = fields
	searchableFieldsMutex.Unlock()

	return fields, nil
}

// Search sets a WHERE clause on the repository to filter results
// by name containing the provided value v.
func (r *Repo[T]) Search(v string) *Repo[T] {

	ls, err := getSearchableFields[T]()
	if err != nil {
		panic(fmt.Errorf("error getting string keys from JSON: %v", err))
	}

	v = "%" + v + "%"
	for i, key := range ls {
		if i == 0 {
			r.DB = r.DB.Where(key+" LIKE ?", v)
		} else {
			r.DB = r.DB.Or(key+" LIKE ?", v)
		}
	}

	return r
}

// Search returns a new Repo instance for type T with a WHERE clause
// filtering by name containing the provided value v.
func Search[T any](v string) *Repo[T] {

	return Do[T]().Search(v)
}
