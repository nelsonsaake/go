package ufs

import (
	"os"
)

func IsDir(fn string) bool {

	fileInfo, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
