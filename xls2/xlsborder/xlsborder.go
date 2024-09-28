package xlsborder

import (
	"github.com/xuri/excelize/v2"
)

func Side(color, side string) excelize.Border {
	return excelize.Border{
		Type:  side,
		Color: color,
		Style: 1,
	}
}

func All(color string) []excelize.Border {
	return []excelize.Border{
		Side(color, "top"),
		Side(color, "left"),
		Side(color, "bottom"),
		Side(color, "right"),
	}
}
