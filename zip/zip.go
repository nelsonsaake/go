package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Zip zips the given files and directories and returns the path to the zip file and error if any
func Zip(zipName string, paths ...string) (string, error) {
	if len(paths) == 0 {
		return "", fmt.Errorf("no paths provided")
	}

	zipPath := zipName
	if filepath.Ext(zipPath) != ".zip" {
		zipPath += ".zip"
	}

	zipPathDir := filepath.Dir(zipPath)
	err := os.MkdirAll(zipPathDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	zipFile, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, p := range paths {
		absPath, err := filepath.Abs(p)
		if err != nil {
			return "", err
		}
		info, err := os.Stat(absPath)
		if err != nil {
			return "", err
		}
		if info.IsDir() {
			err = zipDir(absPath, zipWriter)
			if err != nil {
				return "", err
			}
		} else {
			err = zipFileFunc(absPath, zipWriter)
			if err != nil {
				return "", err
			}
		}
	}
	return zipPath, nil
}

// zipDir zips a directory recursively
func zipDir(absPath string, zipWriter *zip.Writer) error {
	return filepath.Walk(absPath, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(filepath.Dir(absPath), file)
		if err != nil {
			return err
		}
		if info.IsDir() {
			if relPath == "." {
				return nil
			}
			header := &zip.FileHeader{
				Name:   relPath + "/",
				Method: zip.Deflate,
			}
			_, err := zipWriter.CreateHeader(header)
			return err
		}
		return zipFileFunc(file, zipWriter)
	})
}

// zipFileFunc zips a single file
func zipFileFunc(absPath string, zipWriter *zip.Writer) error {
	info, err := os.Stat(absPath)
	if err != nil {
		return err
	}
	relPath := filepath.Base(absPath)
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = relPath
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	fileReader, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer fileReader.Close()
	_, err = io.Copy(writer, fileReader)
	return err
}
