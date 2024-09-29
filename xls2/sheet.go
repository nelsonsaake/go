package xls2

import (
	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	*Doc
	name   string
	styles map[string]excelize.Style
}

func (s *Sheet) SetValue(loc string, value any) error {
	err := s.File.SetCellValue(s.name, loc, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sheet) SetColWidth(col string, width float64) error {
	return s.File.SetColWidth(s.name, col, col, width)
}

func (s *Sheet) SetRowHeight(row int, height float64) error {
	return s.File.SetRowHeight(s.name, row, height)
}

// Println: start from col A, row `row` and each entry in a different cell from left to right
func (s *Sheet) Println(row int, entries ...any) error {

	for i, entry := range entries {

		// col to write of the cell to write entry to
		// row is given as a parameter
		col := i + 1

		// convert, col and row to cell
		cell, err := excelize.CoordinatesToCellName(col, row)
		if err != nil {
			return err
		}

		// write to cell
		err = s.SetValue(cell, entry)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sheet) PageBreak(cell string) error {
	return s.File.InsertPageBreak(s.name, cell)
}
