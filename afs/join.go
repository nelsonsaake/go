package afs

import "path/filepath"

func Join(paths ...string) string {
	return filepath.Join(paths...)
}
