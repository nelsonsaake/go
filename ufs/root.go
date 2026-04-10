package ufs

import (
	"github.com/nelsonsaake/go/afs"
)

func Root() (string, error) {
	return afs.Root(), nil
}
