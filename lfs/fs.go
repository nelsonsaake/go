package lfs

import "github.com/nelsonsaake/go/afs"

type lfs struct{}

func Path(path ...string) string {
	return afs.Path(path...)
}
