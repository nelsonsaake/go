package dto

import (
	"fmt"
)

func transform(raw map[string]any, dtoname string) (map[string]any, error) {

	dto := GetDTO(dtoname)

	var (
		res = map[string]any{}
		err error
	)

	die := func(f string, v ...any) (map[string]any, error) {
		return nil, fmt.Errorf(f, v...)
	}

	// this will copy the map except with
	// the key_case transformed using transformKey
	raw = cloneMap(raw)

	// if inclusion strategy is include all,
	// copy all
	if isIncludeAll() {

		for k, v := range raw {

			// apply case transformation
			k = transformKey(k)

			res[k] = v
		}
	}

	for k, op := range dto {

		k = transformKey(k)
		v, vExists := raw[k]

		switch op := op.(type) {
		case includeValue:
			res[k] = v
		case excludeValue:
			delete(res, k)
		case transformValue:
			if vExists {
				res[k], err = GetFunc(op.name)(v)
				if err != nil {
					return die("error transforming %v: %v", k, err)
				}
			}
		case transformType:
			if vExists {
				res[k], err = Transform(v, op.dtoname)
				if err != nil {
					return die("error transforming %v: %v", k, err)
				}
			}
		default:
			if isIncludeAll() {
				res[k] = op
			}
		}
	}

	return res, nil
}

func transformv(v any, dtoname string) (map[string]any, error) {

	var (
		raw = map[string]any{}
		err error
	)

	die := func(f string, v ...any) (map[string]any, error) {
		return nil, fmt.Errorf(f, v...)
	}

	err = cast(v, &raw)
	if err != nil {
		return die("error casting input to map[string]any: %w", err)
	}

	return transform(raw, dtoname)
}
