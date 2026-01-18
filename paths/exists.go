package paths

import (
	"os"
	"path/filepath"
)

// Exists: checks if file exists from root, if root is found
// or as is
func (p *pfs) Exists(ls ...string) bool {

	var relPath = filepath.Join(ls...)
	if Exists(relPath) {
		return true
	}

	var absPath = filepath.Join(p.root, relPath)
	if Exists(absPath) {
		return true
	}

	return false
}

// Exists returns whether the given file or directory Exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
