package local

import (
	"fmt"
	"path/filepath"
	"strings"

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
		rel      = l.relPath(path)
	)

	err = ufs.WriteFile(path, string(src))
	if err != nil {
		return die("error writing file: %w", err)
	}

	return rel, nil
}

func (l local) relPath(path string) string {

	return strings.TrimPrefix(
		ufs.CleanPath(path),
		ufs.CleanPath(l.dir),
	)
}

func (l local) absPath(path string) string {

	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(
		ufs.CleanPath(l.dir),
		ufs.CleanPath(path),
	)
}

func (l local) Delete(path string) (err error) {

	abs := l.absPath(path)

	return ufs.DelFile(abs)
}

func New(dir string) Storage {
	return &local{dir: dir}
}
