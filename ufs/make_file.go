package ufs

import (
	"io/fs"
	"os"
)

func MakeFile(path string) error {
	return os.WriteFile(path, []byte(""), fs.ModePerm.Perm())
}
