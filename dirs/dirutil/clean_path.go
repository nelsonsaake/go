package dirutil

import "path/filepath"

func CleanPath(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}
