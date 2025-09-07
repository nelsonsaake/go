package dto

import (
	"fmt"
)

func Transformls(v any, dtoname string) ([]map[string]any, error) {

	var (
		raw = []map[string]any{}
		res = []map[string]any{}
		err error
	)

	die := func(f string, v ...any) ([]map[string]any, error) {
		return nil, fmt.Errorf(f, v...)
	}

	err = cast(v, &raw)
	if err != nil {
		return die("error casting input to map[string]any: %w", err)
	}

	for i, v := range raw {

		transformedv, err := transform(v, dtoname)
		if err != nil {
			return die("error casting [%v]: %v", i, err)
		}

		res = append(res, transformedv)
	}

	return res, nil
}
