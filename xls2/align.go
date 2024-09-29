package xls2

import "github.com/xuri/excelize/v2"

var (
	AlignCenter = &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
		Indent:     1,
		WrapText:   true,
	}

	AlignCenterRight = &excelize.Alignment{
		Horizontal: "right",
		Vertical:   "center",
		Indent:     1,
		WrapText:   true,
	}

	AlignCenterLeft = &excelize.Alignment{
		Horizontal: "left",
		Vertical:   "center",
		Indent:     1,
		WrapText:   true,
	}
)
