package ufs

import (
	"fmt"
	"os"
)

func Rename(from, to string) error {

	exists, err := Exists(from)
	if !exists {
		return fmt.Errorf("from(source) file does not exist")
	}

	if err != nil {
		return fmt.Errorf("error checking if from(source) file exists: %v", err)
	}

	exists, err = Exists(to)
	if exists {
		return fmt.Errorf("to(dest) file already exist")
	}

	if err != nil {
		return fmt.Errorf("error checking if to(dest) file exists: %v", err)
	}

	return os.Rename(from, to)
}
