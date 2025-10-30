package tmpl

import "path/filepath"

func JoinFunc(elem ...string) string {
	return filepath.ToSlash(filepath.Clean(filepath.Join(elem...)))
}
