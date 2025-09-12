package logs

import (
	"log"
)

func NoTimestamp() {
	log.SetFlags(0) // no timestamp, no file info
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
}
