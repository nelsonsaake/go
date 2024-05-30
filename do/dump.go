package do

import (
	"encoding/json"
	"fmt"
	"projects/semper-server/pkg/ufs"
	"sync"
)

func Dump(loc string, sm sync.Map) error {

	m := map[string]any{}

	sm.Range(
		func(key, value any) bool {
			m[fmt.Sprint(key)] = value
			return true
		},
	)

	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	err = ufs.WriteFile(loc, string(b))
	if err != nil {
		return err
	}

	return nil
}
