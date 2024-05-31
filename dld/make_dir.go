package dld

import (
	"io/fs"
	"os"
)

// makeDir: make dir with all it's parents
func makeDir(dir string) (err error) {
	return os.MkdirAll(dir, fs.ModePerm.Perm())
}
