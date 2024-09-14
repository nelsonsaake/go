package ufs	

import (
    "path/filepath"
    "strings"
)

func GetFileName(path string) string {

	// Extract the base name (e.g., "myfile.txt")
	return filepath.Base(path) 
}