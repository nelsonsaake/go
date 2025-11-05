package servers

import "strings"

func cleanPort(port string) string {

	port = strings.TrimSpace(port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
