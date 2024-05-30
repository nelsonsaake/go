package ufs

import (
	"os"
	"path/filepath"
)

func isgomodpath(d string) bool {
	return IsFile(filepath.Join(d, "go.mod"))
}

func ModDir() (string, error) {

	badDir := ""

	wd, err := os.Getwd()
	if err != nil {
		return badDir, ErrGettingCurrentDir
	}

	for {
		if isgomodpath(wd) {
			return wd, nil
		}

		if wd == filepath.Dir(wd) {
			return badDir, ErrModDirNotFound
		}

		wd = filepath.Dir(wd)
	}
}
