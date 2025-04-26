package objs

import (
	"encoding/json"
	"fmt"
	"os"
)

func FromFile(path string) (*Obj, error) {

	die := func(f string, a ...any) (*Obj, error) {
		return nil, fmt.Errorf(f, a...)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return die("error reading file: %v", err)
	}

	res := map[string]any{}

	err = json.Unmarshal(raw, &res)
	if err != nil {
		return die("error unmarshalling file content: %v", err)
	}

	return FromMap(res), nil
}
