package ufs

import (
	"os"
)

func IsFile(fn string) bool {

	fileInfo, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		// fmt.Print("\n\n\n", "Is File Err: ", err, "\n\n\n")
		return false
	}

	return !fileInfo.IsDir() // if it is not dir; then is file
}
