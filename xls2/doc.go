package xls2

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Doc struct {
	File   *excelize.File
	Sheets []*Sheet
}

func New() *Doc {
	return &Doc{excelize.NewFile(), []*Sheet{}}
}

func (doc *Doc) NewSheet(name string) *Sheet {

	sheet1 := "Sheet1"

	if doc.File.SheetCount == 1 && doc.File.GetSheetName(0) == sheet1 {
		doc.File.SetSheetName(sheet1, name)
	} else {
		doc.File.NewSheet(name)
	}

	var newSheet = &Sheet{
		doc,
		name,
		map[string]excelize.Style{},
	}
	doc.Sheets = append(doc.Sheets, newSheet)

	return newSheet
}

func (doc *Doc) SaveAs(name string) error {

	for _, sheet := range doc.Sheets {
		for loc, style := range sheet.styles {

			if !(IsCell(loc) || IsRange(loc)) {
				continue
			}

			var (
				tl  = loc // top left cell
				br  = loc // bottom right cell
				err error
			)
			if IsRange(loc) {
				tl, br, err = SplitRange(loc)
				if err != nil {
					continue
				}
			}

			styleID, err := doc.File.NewStyle(&style)
			if err != nil {
				return fmt.Errorf("error creating new style: %v", style)
			}

			err = doc.File.SetCellStyle(sheet.name, tl, br, styleID)
			if err != nil {
				return fmt.Errorf("error setting style: %v", err)
			}

			print(tl, br)
		}
	}

	return doc.File.SaveAs(name)
}
