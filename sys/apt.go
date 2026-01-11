package sys

import (
	"fmt"
)

func AptInstall(pkg ...string) error {

	apt := "apt-get"
	run := NI().Run
	die := fmt.Errorf

	if !IsInstalled(apt) {
		return die("%s not found", apt)
	}

	out, err := run(apt, "update")
	if err != nil {
		return die("%s update failed: %v: %s", apt, err, out)
	}

	out, err = run(apt, "install", "-y", pkg)
	if err != nil {
		return die("%s install %v failed: %v: %s", apt, pkg, err, out)
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
