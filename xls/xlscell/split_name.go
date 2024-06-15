package xlscell

import "github.com/xuri/excelize/v2"

// SplitName: split cell Name into col and row
func SplitName(cell string) (col string, row int, err error) {
	return excelize.SplitCellName(cell)
}
