package do

func LocalhostWithPort(port string) string {
	return "http://localhost" + CleanPort(port)
}
