package styles

import "github.com/xuri/excelize/v2"

var (
	Center = &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
		Indent:     1,
		WrapText:   true,
	}

	CenterRight = &excelize.Alignment{
		Horizontal: "right",
		Vertical:   "center",
		Indent:     1,
		WrapText:   true,
	}

	CenterLeft = &excelize.Alignment{
		Horizontal: "left",
		Vertical:   "center",
		Indent:     1,
		WrapText:   true,
	}
)
