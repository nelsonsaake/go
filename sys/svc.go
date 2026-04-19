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

func systemctl(args ...any) error {

	err := LookPath("systemctl").Error
	if err != nil {
		return fmt.Errorf("error doing lookpath for systemctl: %v", err)
	}

	return Command("systemctl", args...).Run().Error
}

// RELOAD DAEMON COMMAND

func SvcReloadDaemon() error {

	err := systemctl("daemon-reload")
	if err != nil {
		return fmt.Errorf("error reloading daemon: %v", err)
	}

	return nil
}

// ENABLE SERVICE COMMAND

func SvcEnable(svc string) error {

	err := systemctl("enable", "--now", svc)
	if err != nil {
		return fmt.Errorf("error enabling %s: %v", svc, err)
	}

	return nil
}

func SvcStart(svc string) error {

	err := systemctl("start", "--now", svc)
	if err != nil {
		return fmt.Errorf("error starting %s: %v", svc, err)
	}

	return nil
}

func SvcRestart(svc string) error {

	err := systemctl("restart", "--now", svc)
	if err != nil {
		return fmt.Errorf("error restarting %s: %v", svc, err)
	}

	return nil
}

func SvcStop(svc string) error {

	err := systemctl("stop", "--now", svc)
	if err != nil {
		return fmt.Errorf("error stopping %s: %v", svc, err)
	}

	return nil
}

func SvcDisable(svc string) error {

	err := systemctl("disable", "--now", svc)
	if err != nil {
		return fmt.Errorf("error disabling %s: %v", svc, err)
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
