package kernel

var verbose bool

func SetVerbose(v bool) {
	verbose = v
}

func IsVerbose() bool {
	return verbose
}
