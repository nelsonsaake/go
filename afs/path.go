package afs

import (
	"fmt"
	"os"
	"path/filepath"
)

// Path: resolve from the root of the app, if we find the root
// otherwise, we return it as it
func Path(path string) (string, error) {

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a)
	}

	// check if file exists
	_, err := os.Stat("path")
	if err != nil {
		return die("%v", err)
	}

	// find the root of the app
	cwd, err := Root()
	if err != nil {
		// return path as is
		return path, nil
	}

	// path from the root of the app
	path = filepath.Join(cwd, path)

	return path, nil
}
