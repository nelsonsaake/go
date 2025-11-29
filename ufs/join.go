package ufs

import "path/filepath"

func Join(elems ...string) string {
	return filepath.Join(elems...)
}
