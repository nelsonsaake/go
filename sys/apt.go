package sys

import (
	"fmt"
)

func AptInstall(pkg ...string) error {

	apt := "apt-get"
	die := fmt.Errorf

	if !IsInstalled(apt) {
		return die("%s not found", apt)
	}

	err := Command(apt, "update").NI().Run()
	if err != nil {
		return die("%s update failed: %v", apt, err)
	}

	err = Command(apt, "install", "-y", pkg).NI().Run()
	if err != nil {
		return die("%s install %v failed: %v", apt, pkg, err)
	}

	return nil
}

func AptInstallAny(pkg ...string) error {

	var lastErr error
	for _, p := range pkg {

		err := AptInstall(p)
		if err == nil {
			return nil
		}
		lastErr = err
	}

	return lastErr
}
