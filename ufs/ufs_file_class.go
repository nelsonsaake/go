package ufs

import "path/filepath"

func FileClass(fileName string) (class string, err error) {

	switch filepath.Ext(fileName) {
	case jpg, jpeg, png:
		class = ClassImg
	case xls, xlxs, xlw, xlsm, xlsb, xlt, xlr:
		class = ClassXls
	case zip:
		class = ClassZip
	default:
		return "", ErrUnknownFileType
	}

	return
}
