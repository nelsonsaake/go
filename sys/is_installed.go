package sys

import (
	"os/exec"
	"slices"
	"strings"
)

func IsExecLookupPath(path string) bool {

	_, err := exec.LookPath(path)
	return err == nil
}

func IsDpkgPackageInstalled(pkg string) (bool, error) {

	var dump string

	err := Command("dpkg-query", "-W", "-f=${Status}", pkg).WithDump(&dump).Run()
	if err != nil {
		return false, err
	}

	if strings.Contains(dump, "install ok installed") {
		return true, nil
	}

	return false, nil
}

func IsInstalled(path string) bool {

	_, err := exec.LookPath(path)
	if IsExecLookupPath(path) {
		return true
	}

	_, err = IsDpkgPackageInstalled(path)
	if err == nil {
		return true
	}

	return false
}

func IsAnyInstalled(paths ...string) bool {
	return slices.ContainsFunc(paths, IsInstalled)
}
