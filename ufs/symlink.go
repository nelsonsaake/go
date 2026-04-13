package ufs

import (
	"fmt"
	"os"
)

func IsSymlinkPath(path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return false, err
	}
	if err != nil {
		return false, nil
	}
	return info.Mode()&os.ModeSymlink != 0, nil
}

func CreateSymlink(oldname string, newname string) error {
	return os.Symlink(oldname, newname)
}

func IsSymlinkMatch(oldname string, newname string) (bool, error) {

	die := fmt.Errorf

	isSymlink, err := IsSymlinkPath(newname)
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

func EnsureSymlink(oldname string, newname string) error {

	die := fmt.Errorf

	// check if symlink exists
	isSymlinkPointingToTarget, err := IsSymlinkMatch(oldname, newname)
	if err != nil {
		return die("error checking if link if %s -> %s: %w", newname, oldname, err)
	}

	isSysmlinkFileExists, err := IsSymlinkPath(newname)
	if err != nil {
		return die("error checking if symlink exists %s: %w", newname, err)
	}

	// if symlink exsits but is not pointing to the target, remove it
	if isSysmlinkFileExists && !isSymlinkPointingToTarget {
		err := RemoveSymlink(newname)
		if err != nil {
			return die("error removing existing file at %s: %w", newname, err)
		}
	}

	// if symlink does not exist, create it
	if !isSysmlinkFileExists || !isSymlinkPointingToTarget {

		// create service enable symlink if missing
		err := CreateSymlink(oldname, newname)
		if err != nil {
			return die("error creating symlink %s -> %s: %w", newname, oldname, err)
		}
	}

	return nil
}
