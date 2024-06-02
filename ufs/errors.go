package ufs

import (
	"fmt"
)

var (
	ErrUnknownFileType   = fmt.Errorf("unknow file type")
	ErrModDirNotFound    = fmt.Errorf("mod file not found")
	ErrGettingCurrentDir = fmt.Errorf("error getting current dir")
)
