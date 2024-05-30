package do

import (
	"os"
	"projects/semper-server/pkg/str"
	"strings"
)

func GetPort(defaultPort string) (port string) {

	port = os.Getenv("PORT")
	if str.Empty(port) {
		port = defaultPort
	}

	port = strings.TrimSpace(port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
