package ufs

import (
	"fmt"
	"os"
)

func SymlinkExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Symlink(oldname string, newname string) error {
	return os.Symlink(oldname, newname)
}

func IsSymlink(oldname string, newname string) (bool, error) {

	die := fmt.Errorf

	isSymlink, err := SymlinkExists(newname)
	if err != nil {
		return false, die("error checking if symlink exists: %v", err)
	}

	if !isSymlink {
		return false, nil
	}

	target, err := os.Readlink(newname)
	if err != nil {
		return false, die("error reading symlink: %v", err)
	}

	isExists, err := Exists(target)
	if err != nil {
		return false, die("error checking if target exists: %v", err)
	}

	if !isExists {
		return false, nil
	}

	targetAbs, err := Abs(target)
	if err != nil {
		return false, die("error getting absolute path of target: %v", err)
	}

	newnameAbs, err := Abs(newname)
	if err != nil {
		return false, die("error getting absolute path of newname: %v", err)
	}

	if targetAbs != newnameAbs {
		return false, nil
	}

	return true, nil
}

func RemoveSymlink(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
