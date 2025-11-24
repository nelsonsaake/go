package repo

import (
	"fmt"

	"github.com/nelsonsaake/go/fctx"
	"github.com/nelsonsaake/go/strs"
)

type GetRequest = fctx.Request

type GetResponse[T any] = fctx.Paginated[T]

// newResponse creates a new GetResponse for the given request,
// including pagination metadata.
func (r *Repo[T]) newResponse(arg GetRequest) (*GetResponse[T], error) {

	count, err := r.Count()
	if err != nil {
		return nil,
			fmt.Errorf(
				"error counting %q: %w",
				strs.ToWords(r.tables),
				err,
			)
	}

	pageSize := int64(arg.PageSize)
	if pageSize <= 0 {
		pageSize = 1
	}

	res := &GetResponse[T]{
		Request:    arg,
		Data:       make([]T, 0),
		Total:      count,
		TotalPages: (count + pageSize - 1) / pageSize,
	}

	return res, nil
}

// Get retrieves records of type T based on the provided GetRequest,
// supporting search, relations, and pagination.
func (r *Repo[T]) Get(arg GetRequest) (*GetResponse[T], error) {

	res, err := r.newResponse(arg)
	if err != nil {
		return nil, err
	}

	if arg.Search != "" {
		r = r.Search(arg.Search)
	}

	if len(arg.With) > 0 {
		r.With(arg.With...)
	}

	var items = make([]T, 0)

	if arg.NoPagination {
		items, err = r.GetAll()
	} else {
		items, err = r.Paginate(arg.Page, arg.PageSize)
	}

	if err != nil {
		return nil, err
	}

	res.Data = items

	return res, nil
}

// Get is a convenience function that retrieves records of
// type T using the default repository instance.
func Get[T any](arg GetRequest) (*GetResponse[T], error) {

	return Do[T]().Get(arg)
}
