package ufs

import "path/filepath"

func FileCategory(fileName string) (class string, err error) {

	switch filepath.Ext(fileName) {
	case jpg, jpeg, png:
		class = CategoryImage
	case xls, xlxs, xlw, xlsm, xlsb, xlt, xlr:
		class = CategoryXLS
	case zip:
		class = CategoryArchive
	default:
		return "", ErrUnknownFileType
	}

	return
}
