package tfm

import "path/filepath"

func FuncJoin(elem ...string) string {
	return filepath.ToSlash(filepath.Clean(filepath.Join(elem...)))
}
