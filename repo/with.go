package repo

import "github.com/nelsonsaake/go/strs"

// With preloads the specified relations on the repository instance,
// allowing eager loading of related data.
func (r *Repo[T]) With(rel ...string) *Repo[T] {

	for _, v := range rel {
		v = strs.ToPascalCase(v)
		r.DB = r.DB.Preload(v)
	}

	return r
}

// With returns a new Repo instance for type T with the specified relations
// preloaded for eager loading.
func With[T any](rel ...string) *Repo[T] {

	return Do[T]().With(rel...)
}
