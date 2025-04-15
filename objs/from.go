package objs

import (
	"fmt"
)

func From(from any) (*Map, error) {

	to := map[string]any{}

	die := func(f string, a ...any) (*Map, error) {
		return nil, fmt.Errorf(f, a...)
	}

	err := cast(from, &to)
	if err != nil {
		return die("error casting input: %v", err)
	}

	return FromMap(to), nil
}
