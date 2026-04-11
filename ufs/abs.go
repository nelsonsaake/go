package ufs

import "path/filepath"

func Abs(path string) (string, error) {
	return filepath.Abs(path)
}
