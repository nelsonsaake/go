package dirutil

import "path/filepath"

func JoinPath(elem ...string) string {
	return filepath.ToSlash(filepath.Clean(filepath.Join(elem...)))
}
