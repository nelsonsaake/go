package afs

import (
	"fmt"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

func Root() (string, error) {

	death := func(f string, err error) (string, error) {
		return "", fmt.Errorf(f, err)
	}

	// Load package information
	cfg := &packages.Config{
		Mode: packages.NeedModule, // Only need module information
	}

	pkgs, err := packages.Load(cfg, ".") // Current directory
	if err != nil {
		return death("failed to load package: %v", err)
	}

	// Retrieve the module directory and get the absolute path
	for _, pkg := range pkgs {
		if pkg.Module != nil {
			moduleDir := pkg.Module.Dir
			absDir, err := filepath.Abs(moduleDir)
			if err != nil {
				return death("failed to get absolute path: %v", err)
			}
			return absDir, nil
		}
	}

	return death("module dir not found", nil)
}
