package afs

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

var root string

func SetRoot(path string) {
	root = path
}

func Root() string {
	return root
}

func init() {

	var err error

	root, err = findRoot()
	if err != nil {
		panic("failed to determine root path: " + err.Error())
	}
}

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

func cwd() (string, error) {
	return os.Getwd()
}

// findRoot returns the project root in dev mode, or executable dir in production.
func findRoot() (string, error) {
	p, err := devModeRoot()
	if err == nil {
		return p, nil
	}
	return cwd()
}
