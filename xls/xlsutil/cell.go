package xlsutil

import "fmt"

func Cellx(col string, row int) string {
	return fmt.Sprintf("%s%d", col, row)
}
