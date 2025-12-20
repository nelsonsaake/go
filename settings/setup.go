package settings

import (
	"fmt"
	"os"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/objs"
	"github.com/nelsonsaake/go/ufs"
)

var (
	filePath string
	obj      *objs.Obj
)

type Config struct {
	FilePath string
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

	file, err := os.Create(path)
	if err != nil {
		return die("error creating file: %v", err)
	}

	err = file.Close()
	if err != nil {
		return die("error closing file: %v", err)
	}

	return nil
}

func Setup(c Config) {

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
