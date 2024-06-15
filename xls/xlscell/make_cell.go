package xlscell

import (
	"github.com/xuri/excelize/v2"
)

// MakeCell: makes a cell from col and row
func MakeCell(col string, row int) (string, error) {
	return excelize.JoinCellName(col, row)
}
