package local

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nelsonsaake/go/b64"
	"github.com/nelsonsaake/go/strs"
	"github.com/nelsonsaake/go/ufs"
)

type Storage struct {
	dir string
}

// Save: write the content to a dir, and returns the path or an error if any
func (s Storage) Save(b64content string) (string, error) {

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a...)
	}

	if strs.IsEmpty(b64content) {
		return die("content is empty")
	}

	ext, err := b64.Ext(b64content)
	if err != nil {
		return die("error getting ext from base64 string: %w", err)
	}

	src, err := b64.Decode(b64content)
	if err != nil {
		return die("error converting breaking down base64 string: %w", err)
	}

	var (
		fileName = strs.UUID() + ext
		path     = filepath.Join(s.dir, fileName)
		rel      = s.relPath(path)
	)

	err = ufs.WriteFile(path, string(src))
	if err != nil {
		return die("error writing file: %w", err)
	}

	return rel, nil
}

func (s Storage) SaveAs(path, b64content string) error {

	die := func(f string, a ...any) error {
		return fmt.Errorf(f, a...)
	}

	if strs.IsEmpty(b64content) {
		return die("content is empty")
	}

	ext, err := b64.Ext(b64content)
	if err != nil {
		return die("error getting ext from base64 string: %w", err)
	}

	if !strings.HasSuffix(path, ext) {
		path = strings.TrimSuffix(path, "/")
		path += ext
	}

	src, err := b64.Decode(b64content)
	if err != nil {
		return die("error converting breaking down base64 string: %w", err)
	}

	err = ufs.WriteFile(path, string(src))
	if err != nil {
		return die("error writing file: %w", err)
	}

	return nil
}

func (s Storage) relPath(path string) string {

	return strings.TrimPrefix(
		ufs.CleanPath(path),
		ufs.CleanPath(s.dir),
	)
}

func (s Storage) absPath(path string) string {

	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(
		ufs.CleanPath(s.dir),
		ufs.CleanPath(path),
	)
}

func (s Storage) Delete(path string) (err error) {

	abs := s.absPath(path)

	return ufs.DelFile(abs)
}

func New(dir string) *Storage {
	return &Storage{dir: dir}
}
