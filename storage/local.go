package storage

import (
	"fmt"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/storage/local"
)

func Local() Storage {
	var path = afs.Path("/storage")
	fmt.Println("Local storage path:", path)
	return local.New(path)
}

func Store(file string) (string, error) {
	return Local().Save(file)
}
