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

	dump, err := Command("dpkg-query", "-W", "-f=${Status}", pkg).Runo()
	if err != nil {
		return false, err
	}

	if strings.Contains(dump, "install ok installed") {
		return true, nil
	}

	return false, nil
}

func IsInstalled(path string) bool {

	if IsExecLookupPath(path) {
		return true
	}

	installed, _ := IsDpkgPackageInstalled(path)

	return installed
}

func IsAnyInstalled(paths ...string) bool {
	return slices.ContainsFunc(paths, IsInstalled)
}
