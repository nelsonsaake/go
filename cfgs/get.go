package cfgs

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/objs"
	"github.com/nelsonsaake/go/strs"
	"github.com/nelsonsaake/go/ufs"
)

var cache = sync.Map{}

func _splitKey(arg ...string) (string, string, error) {

	var (
		configDir = "src/configs"
	)

	var (
		keys          = strings.Split(strings.Join(arg, "."), ".")
		pathfragments = slices.Insert(keys, 0, configDir)
	)

	for i := range pathfragments {

		// n: partition i+1
		n := i + 1

		// pfn: path fragments at n
		pfn := pathfragments[:n]

		// p: from the p fragments from the root
		p := afs.Path(pfn...) + ".json"

		// if p doesn't exists, we continue the search to other partitions
		exists, _ := ufs.Exists(p)
		if !exists {
			continue
		}

		return p, strings.Join(keys[i:], "."), nil
	}

	return "", "", fmt.Errorf("config file not found")
}

func splitKey(arg ...string) (string, string, error) {

	k := strings.Join(arg, ".")
	ck := getck("splitKey", k)
	if res, ok := cache.Load(ck); ok {
		ls := res.([]string)
		return ls[0], ls[1], nil
	}

	fk, ok, err := _splitKey(k)
	if err != nil {
		return "", "", err
	}

	ls := []string{fk, ok}
	cache.Store(ck, ls)

	return fk, ok, nil
}

// getfilepath: get file key
func getfilepath(k string) (string, error) {
	fk, _, err := splitKey(k)
	if err != nil {
		return "", fmt.Errorf("error getting file path: %v", err)
	}
	return fk, nil
}

// getik: get inner cfg key
func getik(k string) string {
	_, ok, err := splitKey(k)
	if err != nil {
		panic("error getting ik")
	}
	return ok

}

// isnk: checks if the key is nested
func isnk(k string) bool {
	ik := getik(k)
	return len(ik) > 0
}

// getck: get cache key
func getck(t, k string) string {
	return fmt.Sprintf("%s:%s", t, k)
}

func getCfg(k string) (*objs.Obj, error) {

	die := func(f string, a ...any) (*objs.Obj, error) {
		return nil, fmt.Errorf(f, a...)
	}

	file, err := getfilepath(k)
	if err != nil {
		return die("error getting file key: %v", err)
	}

	cfg, err := objs.FromFile(file)
	if err != nil {
		return die("error loading file: %v", err)
	}

	return cfg, nil
}

func get(k string) (any, error) {

	var (
		cfg *objs.Obj
		res any
	)

	cfg, err := getCfg(k)
	if err != nil {
		return nil, err
	}

	if isnk(k) {
		res = cfg.Get(getik(k))
	} else {
		res = cfg.Data
	}

	return res, nil
}

func Get(arg ...string) any {
	k := strings.Join(arg, ".")
	ck := getck("any", k)
	if res, ok := cache.Load(ck); ok {
		return res
	}

	v, err := get(k)
	if err != nil {
		panic(err)
	}
	cache.Store(ck, v)
	return v
}

func GetBool(arg ...string) bool {
	k := strings.Join(arg, ".")
	key := getck("bool", k)
	if res, ok := cache.Load(key); ok {
		return res.(bool)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetBool(getik(k))
	cache.Store(key, v)
	return v
}

func GetFloat64(arg ...string) float64 {
	k := strings.Join(arg, ".")
	key := getck("float64", k)
	if res, ok := cache.Load(key); ok {
		return res.(float64)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetFloat64(getik(k))
	cache.Store(key, v)
	return v
}

func GetInt(arg ...string) int {
	k := strings.Join(arg, ".")
	key := getck("int", k)
	if res, ok := cache.Load(key); ok {
		return res.(int)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetInt(getik(k))
	cache.Store(key, v)
	return v
}

func GetString(arg ...string) string {
	k := strings.Join(arg, ".")
	key := getck("string", k)
	if res, ok := cache.Load(key); ok {
		return res.(string)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetString(getik(k))
	cache.Store(key, v)
	return v
}

func GetStringMap(arg ...string) map[string]string {
	k := strings.Join(arg, ".")
	ck := getck("string-map", k)
	if res, ok := cache.Load(ck); ok {
		return res.(map[string]string)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetStringMap(getik(k))
	cache.Store(ck, v)
	return v
}

func GetMap(arg ...string) map[string]any {
	k := strings.Join(arg, ".")
	key := getck("map", k)
	if res, ok := cache.Load(key); ok {
		return res.(map[string]any)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetMap(getik(k))
	cache.Store(key, v)
	return v
}

func GetSlice(arg ...string) []any {
	k := strings.Join(arg, ".")
	ck := getck("slice", k)
	if res, ok := cache.Load(ck); ok {
		return res.([]any)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetSlice(getik(k))
	cache.Store(ck, v)
	return v
}

func GetStringSlice(arg ...string) []string {
	k := strings.Join(arg, ".")
	key := getck("stringslice", k)
	if res, ok := cache.Load(key); ok {
		return res.([]string)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}
	v := cfg.GetStringSlice(getik(k))
	cache.Store(key, v)
	return v
}

func GetObj(arg ...string) *objs.Obj {
	k := strings.Join(arg, ".")
	ck := getck("obj", k)
	if res, ok := cache.Load(ck); ok {
		return res.(*objs.Obj)
	}

	cfg, err := getCfg(k)
	if err != nil {
		panic(err)
	}

	if strs.IsEmpty(getik(k)) {
		return cfg
	}

	v := cfg.GetObj(getik(k))
	cache.Store(ck, v)
	return v
}
