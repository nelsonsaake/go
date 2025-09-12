package fctx

type DTOFunc = func(v any, name string) (any, error)

var (
	transform DTOFunc
)

func Transform(v any, dtoname string) (any, error) {

	if transform == nil {
		panic("DTOFunc not set")
	}

	return transform(v, dtoname)
}
