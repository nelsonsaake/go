package ufs

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

// devModeRoot returns the root of the project in development mode.
func devModeRoot() (string, error) {

	cfg := &packages.Config{
		Mode: packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return "", fmt.Errorf("failed to load package: %v", err)
	}

	for _, pkg := range pkgs {

		if pkg.Module == nil {
			continue
		}

		moduleDir := pkg.Module.Dir

		absDir, err := filepath.Abs(moduleDir)
		if err != nil {
			return "", fmt.Errorf("failed to get absolute path: %v", err)
		}

		return absDir, nil
	}

	return "", fmt.Errorf("module dir not found")
}

// productionModeRoot returns the directory containing the executable.
func productionModeRoot() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %v", err)
	}
	return filepath.Dir(exePath), nil
}

// Root returns the project root in dev mode, or executable dir in production.
func Root() (string, error) {
	if os.Getenv("GO_DEV_MODE") == "1" {
		return devModeRoot()
	}
	return productionModeRoot()
}
