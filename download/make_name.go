package download

import (
	"fmt"
	"time"
)

// makeName: generate a makeName base on the current timestamp, and the file extension from the data
func makeName(data []byte) string {
	return fmt.Sprint(time.Now().UnixNano()) + extractExt(data)
}
