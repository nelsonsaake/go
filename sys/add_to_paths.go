package sys

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// AddToSystemPathUnix permanently adds a directory to PATH on Unix systems.
// Writes to /etc/profile.d/custom-path.sh.
//
// Rules:
// - path must exist
// - if already present → no-op
func AddToSystemPathUnix(newPath string) error {
	die := fmt.Errorf

	// validate OS
	switch runtime.GOOS {
	case "linux", "darwin":
	default:
		return die("unix path update not supported on OS: %s", runtime.GOOS)
	}

	// check path exists
	info, err := os.Stat(newPath)
	if err != nil || !info.IsDir() {
		return die("path does not exist: %s", newPath)
	}

	const filePath = "/etc/profile.d/custom-path.sh"

	// check existing
	data, _ := os.ReadFile(filePath)
	if strings.Contains(string(data), newPath) {
		return nil
	}

	content := fmt.Sprintf("export PATH=$PATH:%s\n", newPath)

	cmd := exec.Command("sudo", "tee", filePath)
	cmd.Stdin = strings.NewReader(content)

	if err := cmd.Run(); err != nil {
		return die("failed writing PATH file: %v", err)
	}

	return nil
}

// AddToPath is OS-neutral wrapper.
// Unsupported OS → panic.
func AddToPath(newPath string) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		return AddToSystemPathUnix(newPath)
	default:
		panic(fmt.Errorf("AddToPath not implemented for OS: %s", runtime.GOOS))
	}
}
