package jfs

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) (map[string]any, error) {

	die := func(f string, a ...any) (map[string]any, error) {
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

	return res, nil
}
