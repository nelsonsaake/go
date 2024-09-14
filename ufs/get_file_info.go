package ufs	
// package fileutils

import (
    "path/filepath"
    "strings"
)

// FileInfo holds extracted file information
type FileInfo struct {
    FileName    string
    FileExtension string
}

// GetFileInfo returns the FileName and FileExtension from the provided path
func GetFileInfo(path string) *FileInfo {
    // Extract the base name (e.g., "myfile.txt")
    base := filepath.Base(path)

    // Split the file name and extension
    ext := filepath.Ext(base)
    name := strings.TrimSuffix(base, ext)

    return &FileInfo{
        FileName:    name,
        FileExtension: ext,
    }
} 