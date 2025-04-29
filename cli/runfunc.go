package cli

import "github.com/nelsonsaake/go/pretty"

type RunFunc func(args ...string)

func printArgs(args ...string) {
	pretty.Print(args)
}
