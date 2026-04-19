package ufs

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureDirExists(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", dir, err)
	}
	return nil
}
