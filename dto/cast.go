package dto

import (
	"encoding/json"
	"fmt"
)

func cast(from, to any) (err error) {

	j, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("err doing json convert: err marshalling: %w", err)
	}

	err = json.Unmarshal(j, to)
	if err != nil {
		return fmt.Errorf("err doing json convert: un-marshalling: %w", err)
	}

	return nil
}
