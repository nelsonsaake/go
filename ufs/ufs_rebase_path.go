package ufs

import (
	"path/filepath"
	"strings"
)

func RebasePath(oldpath, oldbase, newbase string) string {
	toslash := filepath.ToSlash
	oldbase = toslash(oldbase)
	newbase = toslash(newbase)
	newpath := oldpath
	newpath = toslash(newpath)
	newpath = strings.TrimPrefix(newpath, oldbase)
	newpath = filepath.Join(newbase, newpath)
	newpath = toslash(newpath)
	return newpath
}
