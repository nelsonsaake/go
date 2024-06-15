package xlscell

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

// Cells: get __cells in range if it is a range
func __cells(hcell, vcell string) (cells []string, err error) {
	var (
		cols = []string{}
		rows = []int{}
	)

	// hcel

	col, row, err := SplitName(hcell)
	if err != nil {
		return
	}

	cols = append(cols, col)
	rows = append(rows, row)

	// vcell

	col, row, err = SplitName(vcell)
	if err != nil {
		return
	}

	cols = append(cols, col)
	rows = append(rows, row)

	// col to int

	var colsInt = []int{}
	for _, colName := range cols {
		colInt := 0
		colInt, err = excelize.ColumnNameToNumber(colName)
		if err != nil {
			return
		}
		colsInt = append(colsInt, colInt)
	}

	//

	cmin, cmax := minmax(colsInt[0], colsInt[1])
	rmin, rmax := minmax(rows[0], rows[1])

	for i := cmin; i <= cmax; i++ {
		for j := rmin; j <= rmax; j++ {

			colName := ""
			colName, err = excelize.ColumnNumberToName(i)
			if err != nil {
				return
			}

			cell, err := MakeCell(colName, j)

			if err != nil {
				return nil, err
			}

			cells = append(cells, cell)
		}
	}

	return
}

// Cells: get cells in range if it is a range or
// returns loc in a string array if it is a cell
// if it two and both are both valid cells,
// it will treat the two as if it was a range
// if it is more than one cell, it will only return an array of valid cells
func Cells(loc ...string) (cells []string, err error) {

	switch len(loc) {
	case 0:
		return nil, errors.New("no location provided")

	case 1:
		loc := loc[0]
		if IsRange(loc) {
			start, end, _ := SplitRange(loc)
			return Cells(start, end)
		}

		if IsCell(loc) {
			return append(cells, loc), nil
		}

		return nil, errors.New("location is invalid")

	case 2:
		start, end := loc[0], loc[1]
		if IsCell(start) && IsCell(end) {
			return __cells(start, end)
		}

		fallthrough

	default:
		for _, loc := range loc {
			if IsCell(loc) {
				cells = append(cells, loc)
			}
		}
	}

	return
}
