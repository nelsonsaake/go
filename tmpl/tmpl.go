package tmpl

import "path/filepath"

const RootDir = "./resources/templates"

func isDirectChildOfRootDir(path string) bool {
	return CleanPath(filepath.Dir(path)) == CleanPath(RootDir)
}
