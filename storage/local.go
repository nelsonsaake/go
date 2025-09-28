package storage

import "github.com/nelsonsaake/go/storage/local"

func Local() Storage {
	return local.New("/storage")
}

func Store(file string) (string, error) {
	return Local().Save(file)
}
