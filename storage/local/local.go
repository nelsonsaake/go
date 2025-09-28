package local

import (
	"fmt"
	"path/filepath"

	"github.com/nelsonsaake/go/b64"
	"github.com/nelsonsaake/go/strs"
	"github.com/nelsonsaake/go/ufs"
)

type local struct {
	dir string
}

// Save: write the content to a dir, and returns the path or an error if any
func (l local) Save(content string) (string, error) {

	die := func(err error) (string, error) {
		return "", err
	}

	ext, err := b64.Ext(content)
	if err != nil {
		return die(fmt.Errorf("error getting ext from base64 string: %w", err))
	}

	src, err := b64.Decode(content)
	if err != nil {
		return die(fmt.Errorf("error converting breaking down base64 string: %w", err))
	}

	var (
		fileName = strs.UUID() + ext
		path     = filepath.Join(l.dir, fileName)
	)

	err = ufs.WriteFile(path, string(src))
	if err != nil {
		return die(fmt.Errorf("error writing file: %w", err))
	}

	return path, nil
}

func (l local) Delete(path string) (err error) {
	return ufs.DelFile(path)
}

func New(dir string) Storage {
	return &local{dir: dir}
}
