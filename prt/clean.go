// prt: allows formatting port, prepending ports to localhost
package prt

import "strings"

// fmt: takes in a port and return is in this format :PORT_NUMBER
func Fmt(port string) string {

	port = strings.TrimSpace(port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
