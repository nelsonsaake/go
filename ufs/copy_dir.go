package ufs

import (
	"fmt"
	"os"
)

func CopyDir(from, to string) error {

	die := func(f string, a ...any) error {
		return fmt.Errorf(f, a...)
	}

	if !IsDir(from) {
		return die("from(source) is not a dir")
	}

	if !IsDir(to) {
		err := MakeDir(to)
		if err != nil {
			return die("error making dir: %v", err)
		}
	}

	entries, err := os.ReadDir(from)
	if err != nil {
		return die("error reading directory: %v", err)
	}

	for _, entry := range entries {
		var (
			name     = entry.Name()
			srcPath  = Join(from, name)
			destPath = Join(to, name)
		)
		if entry.IsDir() {
			err := CopyDir(srcPath, destPath)
			if err != nil {
				return die("error copying dir: %v", err)
			}
		} else {
			err := CopyFile(srcPath, destPath)
			if err != nil {
				return die("error copying file: %v", err)
			}
		}
	}

	return nil
}
