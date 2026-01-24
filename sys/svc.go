package sys

import (
	"fmt"
	"os/exec"
	"slices"
)

func LookPath(cmd string) *CmdResults {
	dump, err := exec.LookPath(cmd)
	return &CmdResults{
		Dump:  dump,
		Error: err,
	}
}

// RELOAD DAEMON COMMAND

func SvcReloadDaemon() error {

	err := LookPath("systemctl").Error
	if err != nil {
		return fmt.Errorf("error doing lookpath for systemctl: %v", err)
	}

	err = Command("systemctl", "daemon-reload").Run().Error
	if err != nil {
		return fmt.Errorf("error reloading daemon: %v", err)
	}

	return nil
}

// ENABLE SERVICE COMMAND

func SvcEnable(svc string) error {

	err := LookPath("systemctl").Error
	if err != nil {
		return fmt.Errorf("error doing lookpath for systemctl: %v", err)
	}

	err = Command("systemctl", "enable", "--now", svc).Run().Error
	if err != nil {
		return fmt.Errorf("error enabling %s: %v", svc, err)
	}

	return nil
}

func SvcEnableAny(svcs ...string) error {

	var errors []error

	for _, svc := range svcs {

		err := SvcEnable(svc)
		if err == nil {
			return nil
		}

		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}

	// we return the first error because the others are likely altenative services
	return errors[0]
}

// IS ACTIVE COMMAND

func SvcIsActive(svc string) bool {
	return OK("systemctl", "is-active", "--quiet", svc)
}

func SvcIsAnyActive(svcs ...string) bool {
	return slices.ContainsFunc(svcs, SvcIsActive)
}
