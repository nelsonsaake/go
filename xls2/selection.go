package xls2

import (
	"github.com/nelsonsaake/go/xls/xlsborder"
	"github.com/nelsonsaake/go/xls/xlscell"
	"github.com/xuri/excelize/v2"
)

type Selection struct {
	sheet *Sheet
	locs  []string
}

// Select cell or range
func (s *Sheet) Select(loc string) *Selection {
	return &Selection{
		sheet: s,
		locs:  []string{loc},
	}
}

// Select cell or range
func (s *Selection) Select(loc string) *Selection {
	return &Selection{
		sheet: s.sheet,
		locs:  append(s.locs, loc),
	}
}

// Set style
func (s *Selection) SetStyle(style excelize.Style) *Selection {

	var (
		allCells = []string{}
	)

	for _, loc := range s.locs {
		cells, err := xlscell.Cells(loc)
		if err != nil {
			return s
		} else {
			allCells = append(allCells, cells...)
		}
	}

	for _, cell := range allCells {
		old, ok := s.sheet.styles[cell]
		if !ok {
			s.sheet.styles[cell] = style
		} else {
			s.sheet.styles[cell] = mergestyle(old, style)
		}
	}

	return s
}

// Set font color
func (s *Selection) SetFontColor(color string) *Selection {
	return s.SetStyle(
		excelize.Style{
			Font: &excelize.Font{
				Color: color,
			},
		},
	)
}

// Set fill color
func (s *Selection) SetFill(color string) *Selection {
	return s.SetStyle(
		excelize.Style{
			Fill: excelize.Fill{
				Type:  "gradient",
				Color: []string{color, color},
			},
		},
	)
}

// Set alignment
func (s *Selection) SetAlignment(alignment *excelize.Alignment) *Selection {
	return s.SetStyle(
		excelize.Style{
			Alignment: alignment,
		},
	)
}

// Set border
func (s *Selection) SetBorder(color string) *Selection {
	return s.SetStyle(
		excelize.Style{
			Border: xlsborder.All(color),
		},
	)
}
