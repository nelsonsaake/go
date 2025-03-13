package ufs

import "path/filepath"

// PathFromRoot: returns `go mod dir` + `path`
// if `go mod dir` path not fround, simply return `path`
func PathFromRoot(path string) string {

	cwd, err := GetRootDir()
	if err == nil {
		path = filepath.Join(cwd, path)
	}

	return path
}
