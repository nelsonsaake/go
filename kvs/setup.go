package kvs

import (
	"fmt"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/objs"
	"github.com/nelsonsaake/go/ufs"
)

var (
	filePath string
	obj      *objs.Obj
	autoSave bool
)

type Config struct {
	FilePath string
	AutoSave bool
}

func createfileIfNotExists(path string) error {

	die := func(f string, a ...any) error {
		return fmt.Errorf(f, a...)
	}

	exists, err := ufs.Exists(path)
	if err != nil {
		return die("error checking if file exists: %v", err)
	}

	if exists {
		return nil
	}

	err = ufs.WriteFile(path, "{}")
	if err != nil {
		return die("error writing file: %v", err)
	}

	return nil
}

func Setup(c Config) {

	autoSave = c.AutoSave

	die := func(f string, a ...any) {
		panic(fmt.Sprintf("error setting up settings: %v", fmt.Sprintf(f, a...)))
	}

	filePath = afs.Path(c.FilePath)
	var err error

	err = createfileIfNotExists(filePath)
	if err != nil {
		die("error creating file if not exists: %v", err)
	}

	obj, err = objs.FromFile(filePath)
	if err != nil {
		die("error loading file: %v", err)
	}
}
