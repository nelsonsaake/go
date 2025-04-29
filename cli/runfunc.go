package cli

import "github.com/nelsonsaake/go/pty"

type RunFunc func(args ...string)

func printArgs(args ...string) {
	pty.Println(args)
}
