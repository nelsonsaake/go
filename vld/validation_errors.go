package vld

import (
	"encoding/json"
	"fmt"
)

type ValidatorErrors struct {
	Errors map[string][]string
}

func (v *ValidatorErrors) Error() string {

	jsonBytes, err := json.MarshalIndent(v.Errors, "", "  ")
	if err == nil {
		return string(jsonBytes)
	}

	return fmt.Sprint(v.Errors)
}
