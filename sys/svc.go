package sys

import (
	"fmt"
	"os/exec"
	"slices"
)

// RELOAD DAEMON COMMAND

func SvcReloadDaemon() error {

	if _, err := exec.LookPath("systemctl"); err != nil {
		return fmt.Errorf("error reloading daemon: systemctl doesnot exists: %v", err)
	}

	err := Run("systemctl", "daemon-reload").Error
	if err != nil {
		return fmt.Errorf("failed to reload daemon: %v", err)
	}

	return nil
}

// ENABLE SERVICE COMMAND

func SvcEnable(svc string) error {

	if _, err := exec.LookPath("systemctl"); err != nil {
		return fmt.Errorf("error enabling service: systemctl doesnot exists: %v", err)
	}

	err := Run("systemctl", "enable", "--now", svc).Error
	if err != nil {
		return fmt.Errorf("failed to enable/start %s: %v", svc, err)
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
	return Ok("systemctl", "is-active", "--quiet", svc)
}

func SvcIsAnyActive(svcs ...string) bool {
	return slices.ContainsFunc(svcs, SvcIsActive)
}
