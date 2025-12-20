package settings

import (
	"fmt"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/objs"
)

var (
	filePath string
	obj      *objs.Obj
)

type Config struct {
	FilePath string
}

func Setup(c Config) {

	filePath = afs.Path(c.FilePath)

	var err error
	obj, err = objs.FromFile(filePath)
	if err != nil {
		panic(fmt.Errorf("error setting up settings: %v", err))
	}
}
