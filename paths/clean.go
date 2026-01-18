package paths

import "path/filepath"

func Clean(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}
