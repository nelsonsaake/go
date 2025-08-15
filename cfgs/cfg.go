package cfgs

import (
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nelsonsaake/go/objs"
)

// Config holds all parsed configs from multiple directories
type Config struct {
	mu   sync.RWMutex
	data map[string]*objs.Obj
}

// New loads all configs from given directories into memory in parallel
func New(dirs ...string) *Config {
	c := &Config{
		data: make(map[string]*objs.Obj),
	}

	var wg sync.WaitGroup
	fileCh := make(chan string, 100)

	// Start workers to parse files
	workerCount := 8
	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileCh {
				obj, err := objs.FromFile(file)
				if err != nil {
					continue
				}
				key := normalizeKey(file)
				c.mu.Lock()
				c.data[key] = obj
				c.mu.Unlock()
			}
		}()
	}

	// Feed files into channel
	for _, dir := range dirs {
		files := walkJSONFiles(dir)
		for _, f := range files {
			fileCh <- f
		}
	}
	close(fileCh)

	wg.Wait()
	return c
}

// normalizeKey converts file path to a config key (without .json)
func normalizeKey(path string) string {
	key := strings.TrimSuffix(path, filepath.Ext(path))
	return strings.ReplaceAll(key, string(filepath.Separator), ".")
}

// walkJSONFiles finds all .json files recursively in a dir
func walkJSONFiles(dir string) []string {
	var files []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), ".json") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

// get retrieves *objs.Obj for a key
func (c *Config) get(key string) *objs.Obj {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// ---------------------- Get methods ----------------------

func (c *Config) Get(key string) any {
	obj := c.get(key)
	if obj == nil {
		return nil
	}
	return obj.Data
}

func (c *Config) GetString(key string) string {
	obj := c.get(key)
	if obj == nil {
		return ""
	}
	return obj.GetString("")
}

func (c *Config) GetInt(key string) int {
	obj := c.get(key)
	if obj == nil {
		return 0
	}
	return obj.GetInt("")
}

func (c *Config) GetBool(key string) bool {
	obj := c.get(key)
	if obj == nil {
		return false
	}
	return obj.GetBool("")
}

func (c *Config) GetFloat64(key string) float64 {
	obj := c.get(key)
	if obj == nil {
		return 0
	}
	return obj.GetFloat64("")
}

func (c *Config) GetMap(key string) map[string]any {
	obj := c.get(key)
	if obj == nil {
		return nil
	}
	return obj.GetMap("")
}

func (c *Config) GetStringMap(key string) map[string]string {
	obj := c.get(key)
	if obj == nil {
		return nil
	}
	return obj.GetStringMap("")
}

func (c *Config) GetSlice(key string) []any {
	obj := c.get(key)
	if obj == nil {
		return nil
	}
	return obj.GetSlice("")
}

func (c *Config) GetStringSlice(key string) []string {
	obj := c.get(key)
	if obj == nil {
		return nil
	}
	return obj.GetStringSlice("")
}

func (c *Config) GetObj(key string) *objs.Obj {
	obj := c.get(key)
	if obj == nil {
		return nil
	}
	return obj
}

func (c *Config) GetAs(dest any, key string) {
	obj := c.get(key)
	if obj == nil {
		return
	}
	obj.GetAs("", dest)
}
