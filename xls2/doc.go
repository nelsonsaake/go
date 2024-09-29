package xls2

import (
	"encoding/json"
	"fmt"
	"strings"

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

	styleIndex := map[string]int{}

	toKey := func(s excelize.Style) string {
		key, err := json.Marshal(s)
		if err != nil {
			return ""
		}
		return strings.ToLower(string(key))
	}

	// we do this to optimize style creation
	// we don't want to `NewStyle` multiple time for the exact same `s`
	// there is a limit of styles we can create
	toId := func(s excelize.Style) (int, error) {

		key := toKey(s)

		id, ok := styleIndex[key]
		if ok {
			return id, nil
		}

		id, err := doc.File.NewStyle(&s)
		if err != nil {
			return 0, err
		}

		styleIndex[key] = id

		return id, nil
	}

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

			styleID, err := toId(style)
			if err != nil {
				return fmt.Errorf("error creating new style: %v", style)
			}

			err = doc.File.SetCellStyle(sheet.name, tl, br, styleID)
			if err != nil {
				return fmt.Errorf("error setting style: %v", err)
			}
		}
	}

	return doc.File.SaveAs(name)
}
