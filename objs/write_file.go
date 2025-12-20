package objs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Saveable interface {
	map[string]any | *Obj
}

func SaveToFile[T Saveable](v T, path string) error {

	die := func(f string, a ...any) error {
		return fmt.Errorf(f, a...)
	}

	var m = map[string]any{}

	if obj, ok := any(v).(*Obj); ok {
		m = obj.Data
	}

	if mv, ok := any(v).(map[string]any); ok {
		m = mv
	}

	raw, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return die("error marshalling obj content: %v", err)
	}

	err = os.WriteFile(path, raw, 0644)
	if err != nil {
		return die("error writing file: %v", err)
	}

	return nil
}
