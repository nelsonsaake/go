package do

import (
	"fmt"
	"time"
)

// MakeFileName: if no ext is provided, add nothing
// else if at least one ext is provided, add the first, ignore the rest
func MakeFileName(ext ...string) string {

	fileName := fmt.Sprint(time.Now().UnixNano())

	if len(ext) > 0 {
		fileName += ext[0]
	}

	return fileName
}
