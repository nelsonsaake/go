package ufs

import (
	"fmt"
	"path/filepath"
	"time"
)

func MakeNameRaw() string {

	fileName := fmt.Sprint(time.Now().UnixNano())

	return fileName
}

// MakeName: takes in a name, strips the extension to generate a new name
func MakeName(fileName string) string {

	extension := filepath.Ext(fileName)

	fileName = MakeNameRaw()

	fileName += extension

	return fileName
}

func MakeDirName() string {
	return MakeNameRaw()
}
