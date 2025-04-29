package objs

import (
	"encoding/json"
	"fmt"
)

func FromBytes(raw []byte) (*Obj, error) {

	res := map[string]any{}

	die := func(f string, a ...any) (*Obj, error) {
		return nil, fmt.Errorf(f, a...)
	}

	err := json.Unmarshal(raw, &res)
	if err != nil {
		return die("error unmarshalling bytes content: %v", err)
	}

	return FromMap(res), nil
}
