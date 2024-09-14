package ufs	

import (
    "path/filepath"
    "strings"
)

func GetFileExtension(path string) string {

	// Extract the base name (e.g., "myfile.txt")
	base := filepath.Base(path)

	// Split the file name and extension
	ext := filepath.Ext(base) 

	return ext
}