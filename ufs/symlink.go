package ufs

import "os"

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

	target, err := os.Readlink(oldname)
	if err != nil {
		return false, err
	}

	_, err = os.Stat(target)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	targetAbs, err := Abs(target)
	if err != nil {
		return false, err
	}

	newnameAbs, err := Abs(newname)
	if err != nil {
		return false, err
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
