package ports

func Localhost(port string) string {
	return "http://localhost" + Clean(port)
}
