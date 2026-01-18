package afs

import (
	"os"
	"path/filepath"
)

// Exists: checks if file exists from root, if root is found
// or as is
func Exists(ls ...string) bool {

	// join path fragments
	path := filepath.Join(ls...)
	pathFromRoot := path

	// find the root of the app
	cwd, err := Root()
	if err == nil {

		// path from the root of the app
		pathFromRoot = filepath.Join(cwd, path)
	}

	// check if file exists
	_, err = os.Stat(pathFromRoot)
	if err == nil {
		return true
	}

	_, err = os.Stat(path)
	return err == nil
}
