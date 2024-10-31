package dirs

import (
	"fmt"
	"path/filepath"

	"github.com/nelsonsaake/go/ufs"
)

type Dir struct {
	path string
}

func New(path string) (*Dir, error) {

	die := func(err error) (*Dir, error) {
		return nil, err
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return die(fmt.Errorf("error convert path into absolute path: %v", err))
	}

	exists, err := ufs.Exists(path)
	if err != nil {
		return die(err)
	}

	if exists {
		return &Dir{path}, nil
	}

	err = ufs.MakeDir(path)
	if err != nil {
		return die(fmt.Errorf("dir doesn't exist, error making dir: %v", err))
	}

	return &Dir{path}, nil
}
