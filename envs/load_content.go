package envs

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	loadedContent = make(map[string]struct{})
	contentMu     sync.Mutex
)

// loadContent loads environment variables from a content string if it hasn't been loaded yet
func loadContent(content string) error {
	contentHash := fmt.Sprintf("%x", sha256.Sum256([]byte(content)))

	contentMu.Lock()
	if _, ok := loadedContent[contentHash]; ok {
		contentMu.Unlock()
		return nil // already loaded, skip
	}
	contentMu.Unlock()

	envMap, err := godotenv.Unmarshal(content)
	if err != nil {
		return fmt.Errorf("error parsing env content: %w", err)
	}

	for key, value := range envMap {
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("error setting env var %s: %w", key, err)
		}
	}

	contentMu.Lock()
	loadedContent[contentHash] = struct{}{}
	contentMu.Unlock()

	return nil
}

// LoadContent loads environment variables from one or more content strings.
// It skips already-loaded content based on content hash.
func LoadContent(contents ...string) error {
	for _, content := range contents {
		content = strings.TrimSpace(content)
		if content == "" {
			continue
		}
		if err := loadContent(content); err != nil {
			return err
		}
	}
	return nil
}
