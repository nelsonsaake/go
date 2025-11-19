package ufs

import "path/filepath"

// PathFromRoot: returns `go mod dir` + `path`
// if `go mod dir` path not fround, simply return `path`
func PathFromRoot(path string) string {

	root, err := Root()
	if err == nil {
		path = filepath.Join(root, path)
	}

	return path
}
