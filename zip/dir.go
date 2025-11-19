package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Dir zips the given directory and returns the path to the zip file and error if any
func Dir(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	zipPath := absPath + ".zip"
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(absPath, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(absPath, file)
		if err != nil {
			return err
		}
		if info.IsDir() {
			if relPath == "." {
				return nil
			}
			// Add directory entry
			header := &zip.FileHeader{
				Name:   relPath + "/",
				Method: zip.Deflate,
			}
			_, err := zipWriter.CreateHeader(header)
			return err
		}
		// Add file entry
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
		fileReader, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileReader.Close()
		_, err = io.Copy(writer, fileReader)
		return err
	})

	if err != nil {
		return "", err
	}
	return zipPath, nil
}
