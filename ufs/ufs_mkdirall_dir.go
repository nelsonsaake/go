package ufs

import (
	"io/fs"
	"os"
)

func MkdirAll(dir string) (err error) {
	return os.MkdirAll(dir, fs.ModePerm.Perm())
}
