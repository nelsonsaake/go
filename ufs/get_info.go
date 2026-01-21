package ufs

// package fileutils

import (
	"path/filepath"
	"strings"
)

// Info holds extracted file information
type Info struct {
	Name      string
	Extension string
}

// GetInfo returns the Name and Extension from the provided path
func GetInfo(path string) *Info {
	// Extract the base name (e.g., "myfile.txt")
	base := filepath.Base(path)

	// Split the file name and extension
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	return &Info{
		Name:      name,
		Extension: ext,
	}
}
