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

// resolve walks down the tree using the given key parts
// and returns the found object and any remaining innerKey
func (c *Config) resolve(ks ...string) (*objs.Obj, string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ks = strings.Split(strings.Join(ks, "."), ".")
	if len(ks) == 0 {
		return nil, ""
	}

	k0 := ks[0]
	rest := ks[1:]

	// If no subkey, check here
	if len(rest) == 1 {
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
		return cc.data, strings.Join(rest, ".")
	}

	// Continue traversal
	return cc.resolve(rest...)
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

func (c *Config) Get(ks ...string) any {
	obj, innerKey := c.resolve(ks...)
	if obj == nil {
		return nil
	}
	if innerKey == "" {
		return obj.Data
	}
	return obj.Get(innerKey)
}

func (c *Config) GetString(ks ...string) string {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return ""
	}
	return obj.GetString(innerKey)
}

func (c *Config) GetInt(ks ...string) int {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return 0
	}
	return obj.GetInt(innerKey)
}

func (c *Config) GetBool(ks ...string) bool {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return false
	}
	return obj.GetBool(innerKey)
}

func (c *Config) GetFloat64(ks ...string) float64 {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return 0
	}
	return obj.GetFloat64(innerKey)
}

func (c *Config) GetMap(ks ...string) map[string]any {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetMap(innerKey)
}

func (c *Config) GetStringMap(ks ...string) map[string]string {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetStringMap(innerKey)
}

func (c *Config) GetSlice(ks ...string) []any {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetSlice(innerKey)
}

func (c *Config) GetStringSlice(ks ...string) []string {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return nil
	}
	return obj.GetStringSlice(innerKey)
}

func (c *Config) GetObj(ks ...string) *objs.Obj {
	obj, innerKey := c.resolve(ks...)
	if obj == nil {
		return nil
	}
	if innerKey == "" {
		return obj
	}
	return obj.GetObj(innerKey)
}

func (c *Config) GetAs(dest any, ks ...string) {
	obj, innerKey := c.resolve(ks...)
	if obj == nil || innerKey == "" {
		return
	}
	obj.GetAs(innerKey, dest)
}
