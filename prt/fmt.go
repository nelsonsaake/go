package prt

import "strings"

// fmt: takes in a port and return is in this format :PORT_NUMBER
// fmt: makes sure that port has a `:` before the port number
// eg. `3000` or `:3000` and this function will make sure it's always comes out as `:3000`
func Fmt(port string) string {

	port = strings.TrimSpace(port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
