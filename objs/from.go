package objs

import (
	"encoding/json"
	"fmt"
)

func From(from any) (*Map, error) {

	to := map[string]any{}

	die := func(f string, a ...any) (*Map, error) {
		return nil, fmt.Errorf("error casting input: %v", fmt.Errorf(f, a...))
	}

	j, err := json.Marshal(from)
	if err != nil {
		return die("err marshalling: %v", err)
	}

	err = json.Unmarshal(j, &to)
	if err != nil {
		return die("un-marshalling: %v", err)
	}

	return FromMap(to), nil
}
