package do

import "strings"

func CleanPort(port string) string {

	port = strings.TrimSpace(port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
