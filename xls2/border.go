package xls2

import (
	"github.com/xuri/excelize/v2"
)

func BorderSide(color, side string) excelize.Border {
	return excelize.Border{
		Type:  side,
		Color: color,
		Style: 1,
	}
}

func Border(color string) []excelize.Border {
	return []excelize.Border{
		BorderSide(color, "top"),
		BorderSide(color, "left"),
		BorderSide(color, "bottom"),
		BorderSide(color, "right"),
	}
}
