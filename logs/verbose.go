package logs

import (
	"log"
	"os"
)

func Verbose() {
	log.SetOutput(os.Stdout)
}
