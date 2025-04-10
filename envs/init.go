package envs

import "fmt"

func _init() error {
	if isLoaded {
		return fmt.Errorf("env is loaded")
	}

	return nil
}
