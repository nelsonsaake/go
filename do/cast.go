package do

import (
	"encoding/json"
	"fmt"
)

// Cast: provides type casting for struct. Allows 2 structs with overlapping json tags be cast from one to the other.
func Cast(from, to any) (err error) {

	j, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("err marshalling: %v", err)
	}

	err = json.Unmarshal(j, to)
	if err != nil {
		return fmt.Errorf("un-marshalling: %v", err)
	}

	return nil
}
