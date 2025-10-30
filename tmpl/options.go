package tmpl

import (
	"fmt"
	"os"
	"path/filepath"
)

func getOptionFiles(dir string) ([]string, error) {

	files := []string{}

	findOptions := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		if !info.IsDir() {
			files = append(files, CleanPath(path))
		}

		return nil
	}

	err := filepath.Walk(dir, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %v", dir, err)
	}

	return files, err
}

type Option struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func Options() ([]Option, error) {

	dirs := map[string]string{}
	options := []Option{}

	findDirs := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		if info.IsDir() && isDirectChildOfRootDir(path) {
			dirs[info.Name()] = path
		}

		return nil
	}

	err := filepath.Walk(RootDir, findDirs)
	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %v", RootDir, err)
	}

	for name, path := range dirs {

		files, err := getOptionFiles(path)
		if err != nil {
			return nil, fmt.Errorf("error walking the path %s: %v", RootDir, err)
		}

		options = append(options, Option{Name: name, Values: files})
	}

	return options, err
}
