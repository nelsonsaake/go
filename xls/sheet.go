package sxl

import (
	"fmt"

	// please don't take this out
	// adding them triggers a process that is required for this package to function them properly
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ifile"
	"github.com/nelsonsaake/go/xls/xlsborder"
	"github.com/nelsonsaake/go/xls/xlscell"

	"github.com/xuri/excelize/v2"
)

type sheet struct {
	*doc
	name   string
	styles map[string]excelize.Style
}

type Sheet struct {
	*sheet
}

func (s *sheet) SetStyle(style excelize.Style, loc ...string) error {

	cells, err := xlscell.Cells(loc...)
	if err != nil {
		return err
	}

	for _, cell := range cells {
		old, ok := s.styles[cell]
		if !ok {
			s.styles[cell] = style
		} else {
			s.styles[cell] = appendStyle(old, style)
		}
	}

	return nil
}

func (s *sheet) SetFontColor(color string, loc ...string) error {
	return s.SetStyle(
		excelize.Style{
			Font: &excelize.Font{
				Color: color,
			},
		},
		loc...,
	)
}

func (s *sheet) SetFill(color string, loc ...string) error {
	return s.SetStyle(
		excelize.Style{
			Fill: excelize.Fill{
				Type:  "gradient",
				Color: []string{color, color},
			},
		},
		loc...,
	)
}

func (s *sheet) SetAlignment(alignment excelize.Alignment, loc ...string) error {

	return s.SetStyle(
		excelize.Style{
			Alignment: &alignment,
		},
		loc...,
	)
}

func (s *sheet) SetBorder(color string, loc ...string) error {
	return s.SetStyle(
		excelize.Style{
			Border: xlsborder.All(color),
		},
		loc...,
	)
}

func (s *sheet) Merge(loc string) error {
	var hCell, vCell, err = xlscell.SplitRange(loc)
	if err != nil {
		return err
	}

	return s.File.MergeCell(s.name, hCell, vCell)
}

func (s *sheet) SetMergeValue(loc string, value any) error {
	var hCell, _, err = xlscell.SplitRange(loc)
	if err != nil {
		return err
	}

	return s.File.SetCellValue(s.name, hCell, value)
}

func (s *sheet) SetValue(loc string, value any) error {
	err := s.File.SetCellValue(s.name, loc, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *sheet) SetColWidth(col string, width float64) error {
	return s.File.SetColWidth(s.name, col, col, width)
}

func (s *sheet) SetRowHeight(row int, height float64) error {
	return s.File.SetRowHeight(s.name, row, height)
}

func (s *sheet) AddPictureWithOffset(loc, url string, offset Offset) error {

	file, err := ifile.New(url)
	if err != nil {
		return err
	}

	// fs, err := xlsutil.NewPictureFormatSheet(file.NewReader())
	// if err != nil {
	// 	return err
	// }

	ext := file.Ext()
	if ext == ".webp" {
		ext = ".png"
	}

	err = s.File.AddPictureFromBytes(
		s.name,
		loc,
		&excelize.Picture{
			Extension: ext,
			File:      file.Bytes,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *sheet) AddPicturef(loc, url, format string) error {

	file, err := ifile.New(url)
	if err != nil {
		return err
	}

	ext := file.Ext()
	if ext == ".webp" {
		ext = ".png"
	}

	err = s.File.AddPictureFromBytes(
		s.name,
		loc,
		// format,
		// file.Name(),
		// ext,
		// file.Bytes,
		&excelize.Picture{
			Extension: ext,
			File:      file.Bytes,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *sheet) AddPicture(loc, url string) error {

	file, err := ifile.New(url)
	if !(err != nil && do.IsImage(file.Bytes)) {

		file, err = ifile.New(url) // try again
		if err != nil {
			return err
		}

		if !do.IsImage(file.Bytes) {
			return fmt.Errorf("content at %v: is not an image", url)
		}
	}

	if !do.IsImage(file.Bytes) {
		file, err = ifile.New(url) // try again
		if err != nil {
			return err
		}
	}

	// fs, err := xlsutil.NewPictureFormatSheet(file.NewReader())
	// if err != nil {
	// 	return err
	// }

	ext := file.Ext()
	if ext == ".webp" {
		ext = ".png"
	}

	err = s.File.AddPictureFromBytes(
		s.name,
		loc,
		// fs.Scale().AsString(),
		// file.Name(),
		// ext,
		// file.Bytes,
		&excelize.Picture{
			Extension: ext,
			File:      file.Bytes,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// func (s *sheet) Freeze(cols, rows int) (err error) {

// 	colName, err := excelize.ColumnNumberToName(cols)
// 	if err != nil {
// 		return err
// 	}

// 	x, err := xlscell.MakeCell(colName, rows)
// 	if err != nil {
// 		return err
// 	}

// 	return s.File.SetPanes(
// 		s.name,
// 		`{
// 			"freeze":true,
// 			"split":false,

// 			"x_split":`+fmt.Sprint(cols)+`,
// 			"y_split":`+fmt.Sprint(rows)+`,

// 			"top_left_cell":"`+x+`",
// 			"active_pane":"bottomRight",

// 			"panes":[{"sqref":"k16", "active_cell":"k16", "pane":"bottomRight"}]
// 		}`,
// 	)
// }
