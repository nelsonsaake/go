package sxl

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type doc struct {
	File   *excelize.File
	Sheets []*sheet
}

type Doc struct {
	*doc
}

func New() Doc {
	return Doc{&doc{excelize.NewFile(), []*sheet{}}}
}

func (doc *doc) NewSheet(name string) Sheet {

	sheet1 := "Sheet1"

	if doc.File.SheetCount == 1 && doc.File.GetSheetName(0) == sheet1 {
		doc.File.SetSheetName(sheet1, name)
	} else {
		doc.File.NewSheet(name)
	}

	var newSheet = &sheet{
		doc,
		name,
		map[string]excelize.Style{},
	}
	doc.Sheets = append(doc.Sheets, newSheet)

	return Sheet{newSheet}
}

func (doc *doc) SaveAs(name string) error {

	for _, sheet := range doc.Sheets {
		for cell, style := range sheet.styles {
			styleID, err := doc.File.NewStyle(&style)
			if err != nil {
				return fmt.Errorf("error creating new style: %v", style)
			}
			err = doc.File.SetCellStyle(sheet.name, cell, cell, styleID)
			if err != nil {
				return fmt.Errorf("error setting style: %v", err)
			}
		}
	}

	return doc.File.SaveAs(name)
}
