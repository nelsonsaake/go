package xls2

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func appendStyle(old, new excelize.Style) excelize.Style {
	if new.Border != nil {
		old.Border = new.Border
	}

	if fmt.Sprint(new.Fill) != fmt.Sprint(excelize.Fill{}) {
		old.Fill = new.Fill
	}

	if new.Font != nil {
		old.Font = new.Font
	}

	if new.Alignment != nil {
		old.Alignment = new.Alignment
	}

	if new.Protection != nil {
		old.Protection = new.Protection
	}

	if new.NumFmt != 0 {
		old.NumFmt = new.NumFmt
	}

	if new.DecimalPlaces != nil {
		old.DecimalPlaces = new.DecimalPlaces
	}

	if new.CustomNumFmt != nil {
		old.CustomNumFmt = new.CustomNumFmt
	}

	old.NegRed = new.NegRed

	return old
}
