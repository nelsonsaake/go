package afs

import (
	"path/filepath"
)

// Path: resolve from the root of the app, if we find the root
// otherwise, we return it as it
func Path(ls ...string) string {

	// join path fragments
	path := filepath.Join(ls...)

	// find the root of the app
	cwd, err := Root()
	if err == nil {

		// path from the root of the app
		path = filepath.Join(cwd, path)
	}

	return path
}
