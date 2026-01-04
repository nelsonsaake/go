package envs

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/ufs"
)

var (
	loadedFiles = make(map[string]struct{})
	mu          sync.Mutex
)

// loadFile loads a single .env file if it hasn't been loaded yet
func loadFile(path string) error {
	mu.Lock()
	if _, ok := loadedFiles[path]; ok {
		mu.Unlock()
		return nil // already loaded, skip
	}
	mu.Unlock()

	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("error loading %s: %w", path, err)
	}

	mu.Lock()
	loadedFiles[path] = struct{}{}
	mu.Unlock()

	return nil
}

// Load loads one or more .env files, propagating errors.
// It skips already-loaded files.
func Load(paths ...string) error {

	if len(paths) == 0 {
		paths = []string{".env"}
	}

	for _, path := range paths {

		fullPath := afs.Path(path)
		if !ufs.IsFile(fullPath) {
			continue // skip non-existing files
		}

		if err := loadFile(fullPath); err != nil {
			return err
		}
	}
	return nil
}
