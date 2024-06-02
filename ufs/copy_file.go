package ufs

import (
	"fmt"
)

func CopyFile(from, to string) error {
	ferr := func(err error) error {
		return fmt.Errorf("%v \nfrom: %s \nto: %s", err, from, to)
	}
	switch {
	// TODO: copy if dir doesn't copy
	case IsDir(from):
		if err := MakeDir(to); err != nil {
			return ferr(fmt.Errorf("error creating dir: %v", err))
		}
	case IsFile(from):
		content, err := ReadFile(from)
		if err != nil {
			return ferr(fmt.Errorf("error copying dir: %v", err))
		}
		WriteFile(to, content)
	default:
		return ferr(fmt.Errorf("from(source) is not a file or dir"))
	}
	return nil
}
