package cin

import (
	"bufio"
	"os"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}
