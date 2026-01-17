package sys

import (
	"os/exec"
	"slices"
)

func IsExecLookupPath(path string) bool {

	_, err := exec.LookPath(path)
	return err == nil
}

func IsDpkgPackageInstalled(pkg string) bool {
	return Command("dpkg-query", "-W", "-f=${Status}", pkg).Run().Contains("install ok installed")
}

func IsInstalled(path string) bool {
	return IsExecLookupPath(path) || IsDpkgPackageInstalled(path)
}

func IsAnyInstalled(paths ...string) bool {
	return slices.ContainsFunc(paths, IsInstalled)
}
