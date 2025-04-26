package objs

import "fmt"

func SliceFrom(a any) ([]*Obj, error) {

	die := func(f string, a ...any) ([]*Obj, error) {
		return nil, fmt.Errorf(f, a...)
	}

	ls := []any{}

	err := cast(a, &ls)
	if err != nil {
		return die("error casting input: %v", err)
	}

	res := []*Obj{}

	for i, v := range ls {

		obj, err := From(v)
		if err != nil {
			return die("error casting entry %v: %v", i, err)
		}

		res = append(res, obj)
	}

	return res, nil
}
