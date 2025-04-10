package envs

import (
	"fmt"

	"github.com/joho/godotenv"
)

var isLoaded = false

func Load(path string) error {

	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("error loading %s: %v", path, err)
	} else {
		isLoaded = true
	}

	return nil
}
