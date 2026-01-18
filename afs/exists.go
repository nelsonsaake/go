package afs

import (
	"os"
	"path/filepath"
)

// Exists: checks if file exists from root, if root is found
// or as is
func Exists(ls ...string) bool {

	var relPath = filepath.Join(ls...)
	if exists(relPath) {
		return true
	}

	var absPath = filepath.Join(root, relPath)
	if exists(absPath) {
		return true
	}

	return false
}

// Exists returns whether the given file or directory Exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
