package storage

import (
	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/storage/local"
	"github.com/nelsonsaake/go/ufs"
)

var prefix = "/storage/"

func Local() Storage {
	return local.New(Root())
}

func Store(file string) (string, error) {
	return Local().Save(file)
}

func Delete(file string) error {
	return Local().Delete(file)
}

func Path(elem ...string) string {
	return afs.Path(append([]string{prefix}, elem...)...)
}

func Root() string {
	return afs.Path(prefix)
}

func MakeDir(elem ...string) (string, error) {
	dir := Path(elem...)
	err := ufs.MakeDir(dir)
	return dir, err
}
