package ufs

import (
	"os"
)

func DelFile(file string) (err error) {
	if IsFile(file) {
		if err := os.Remove(file); err != nil {
			return err
		}
	}
	return nil
}
