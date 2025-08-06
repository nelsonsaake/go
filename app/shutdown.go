package app

import "fmt"

func Shutdown() error {

	for name, teardown := range teardowns() {
		err := teardown.Teardown()
		if err != nil {
			return fmt.Errorf("error teardown up %s: %v", name, err)
		}
	}

	return nil
}
