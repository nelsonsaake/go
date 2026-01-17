package sys

import "runtime"

const (
	windows = "windows"
	linux   = "linux"
	darwin  = "darwin"
	freebsd = "freebsd"
)

func IsWindows() bool {
	return runtime.GOOS == windows
}

func IsDarwin() bool {
	return runtime.GOOS == darwin
}

func IsLinux() bool {
	return runtime.GOOS == linux
}

func IsFreeBSD() bool {
	return runtime.GOOS == freebsd
}
