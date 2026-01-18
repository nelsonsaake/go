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
	return filepath.Join(root, path)
}
