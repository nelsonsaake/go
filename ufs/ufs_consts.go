package ufs

import (
	"fmt"
)

var (
	// errs
	ErrUnknownFileType   = fmt.Errorf("unknow file type")
	ErrModDirNotFound    = fmt.Errorf("mod file not found")
	ErrGettingCurrentDir = fmt.Errorf("error getting current dir")

	// file types
	jpg  = ".jpg"
	jpeg = ".jpeg"
	png  = ".png"

	xls  = "xls"
	xlxs = "xlxs"
	xlw  = ".xlw"
	xlsm = ".xlsm"
	xlsb = ".xlsb"
	xlt  = ".xlt"
	xlr  = ".xlr"

	zip = ".zip"

	// list of file types
	ACCEPTABLE_EXTENSION = []string{
		jpg, jpeg, png,
		xls, xlxs, xlw, xlsm, xlsb, xlt, xlr,
		zip,
	}

	// file classes
	ClassImg = "img"
	ClassXls = "xls"
	ClassZip = "zip"

	// list of file classes
	FILE_CLASSES = []string{
		ClassImg, ClassXls, ClassZip,
	}
)
