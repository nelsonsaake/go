package ufs

import (
	"os"
	"path/filepath"
)

func WriteFile(fpath, content string) (err error) {

	fpath = filepath.Clean(fpath)

	fdir := filepath.Dir(fpath)

	err = MakeDir(fdir)
	if err != nil {
		return err
	}

	err = os.WriteFile(fpath, []byte(content), 0777)

	return err
}
