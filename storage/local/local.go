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

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a...)
	}

	if strs.IsEmpty(content) {
		return die("content is empty")
	}

	ext, err := b64.Ext(content)
	if err != nil {
		return die("error getting ext from base64 string: %w", err)
	}

	src, err := b64.Decode(content)
	if err != nil {
		return die("error converting breaking down base64 string: %w", err)
	}

	var (
		fileName = strs.UUID() + ext
		path     = filepath.Join(l.dir, fileName)
	)

	rel, err := RelativePath(path)
	if err != nil {
		return die("error getting relative path: %w", err)
	}

	err = ufs.WriteFile(path, string(src))
	if err != nil {
		return die("error writing file: %w", err)
	}

	return rel, nil
}

func (l local) absPath(path string) (string, error) {

	if filepath.IsAbs(path) {
		return path, nil
	}

	root, err := Root()
	if err != nil {
		return "", fmt.Errorf("error getting root path: %w", err)
	}

	return filepath.Join(root, path), nil
}

func (l local) Delete(path string) (err error) {

	abs, err := l.absPath(path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %w", err)
	}

	return ufs.DelFile(abs)
}

func New(dir string) Storage {
	return &local{dir: dir}
}
