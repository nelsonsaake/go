package ufs

import (
	"io/fs"
	"os"
)

func MakeDir(path string) (err error) {
	return os.MkdirAll(path, fs.ModePerm.Perm())
}
