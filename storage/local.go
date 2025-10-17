package storage

import (
	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/storage/local"
)

func Local() Storage {
	var path = afs.Path("/storage")
	return local.New(path)
}

func Store(file string) (string, error) {
	return Local().Save(file)
}

func Delete(file string) error {
	return Local().Delete(file)
}
