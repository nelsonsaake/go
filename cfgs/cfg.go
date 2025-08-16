package cfgs

import (
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/objs"
)

// Config holds all parsed configs from multiple directories
type Config struct {
	mu      sync.RWMutex
	configs map[string]*Config
	data    *objs.Obj
}

func (c *Config) set(key string, obj *objs.Obj) {
	c.mu.Lock()
	defer c.mu.Unlock()

	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		// last part → store the object here
		c.configs[parts[0]] = &Config{
			configs: make(map[string]*Config),
			data:    obj,
		}
		return
	}

	// traverse or create sub-config
	k0 := parts[0]
	cc, ok := c.configs[k0]
	if !ok {
		cc = &Config{configs: make(map[string]*Config)}
		c.configs[k0] = cc
	}
	cc.set(strings.Join(parts[1:], "."), obj)
}

func (c *Config) find(key string) (*objs.Obj, string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	parts := strings.Split(key, ".")
	if len(parts) == 0 {
		return nil, ""
	}

	k0 := parts[0]
	rest := strings.Join(parts[1:], ".")

	// If no subkey, check here
	if len(parts) == 1 {
		cc, ok := c.configs[k0]
		if !ok {
			return nil, ""
		}
		return cc.data, ""
	}

	// Still have deeper keys
	cc, ok := c.configs[k0]
	if !ok {
		return nil, ""
	}

	// If cc has data, stop traversal and pass remaining as innerKey
	if cc.data != nil && len(cc.configs) == 0 {
		return cc.data, rest
	}

	// Continue traversal
	return cc.find(rest)
}

// New loads all configs from given directories into memory in parallel
func New(dirs ...string) *Config {
	c := &Config{
		configs: make(map[string]*Config),
	}

	type fileJob struct {
		base string
		path string
	}

	var wg sync.WaitGroup
	fileCh := make(chan fileJob, 100)

	workerCount := 8
	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range fileCh {
				obj, err := objs.FromFile(job.path)
				if err != nil {
					continue
				}
				key := normalizeKey(job.base, job.path)
				c.set(key, obj)
			}
		}()
	}

	for _, dir := range dirs {
		dir = afs.Path(dir)
		files := walkJSONFiles(dir)
		for _, f := range files {
			fileCh <- fileJob{base: dir, path: f}
		}
	}
	close(fileCh)

	wg.Wait()
	return c
}

// normalizeKey converts file path to a config key (without .json)
func normalizeKey(baseDir, path string) string {
	rel, _ := filepath.Rel(baseDir, path)
	rel = strings.TrimSuffix(rel, filepath.Ext(rel))
	return strings.ReplaceAll(rel, string(filepath.Separator), ".")
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

// ---------------------- Get methods ----------------------
func (c *Config) Get(key string) any {
	obj, innerKey := c.find(key)
	if obj == nil {
		return nil
	}
	if innerKey == "" {
		return obj.Data
	}
	return obj.Get(innerKey)
}

func (c *Config) GetString(key string) string {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return ""
	}
	return obj.GetString(innerKey)
}

func (c *Config) GetInt(key string) int {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return 0
	}
	return obj.GetInt(innerKey)
}

func (c *Config) GetBool(key string) bool {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return false
	}
	return obj.GetBool(innerKey)
}

func (c *Config) GetFloat64(key string) float64 {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return 0
	}
	return obj.GetFloat64(innerKey)
}

func (c *Config) GetMap(key string) map[string]any {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetMap(innerKey)
}

func (c *Config) GetStringMap(key string) map[string]string {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetStringMap(innerKey)
}

func (c *Config) GetSlice(key string) []any {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetSlice(innerKey)
}

func (c *Config) GetStringSlice(key string) []string {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetStringSlice(innerKey)
}

func (c *Config) GetObj(key string) *objs.Obj {
	obj, innerKey := c.find(key)
	if obj == nil {
		return nil
	}
	if innerKey == "" {
		return obj
	}
	return obj.GetObj(innerKey)
}

func (c *Config) GetAs(dest any, key string) {
	obj, innerKey := c.find(key)
	if obj == nil || innerKey == "" {
		return
	}
	obj.GetAs(innerKey, dest)
}
