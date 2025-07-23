package cfgs

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/objs"
	"github.com/nelsonsaake/go/strs"
	"github.com/nelsonsaake/go/ufs"
)

type cfg struct {
	dir   string
	cache sync.Map
}

func New(dir ...string) *cfg {
	return &cfg{
		dir:   filepath.Join(dir...),
		cache: sync.Map{},
	}
}

func (c cfg) _splitKey(arg ...string) (string, string, error) {

	s := strings.Split
	j := strings.Join

	var (
		keys          = s(j(s(j(arg, "."), "."), "/"), "/")
		pathfragments = slices.Insert(keys, 0, c.dir)
	)

	fmt.Println(pathfragments)

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

func (c cfg) splitKey(arg ...string) (string, string, error) {

	k := strings.Join(arg, ".")
	ck := getck("splitKey", k)
	if res, ok := c.cache.Load(ck); ok {
		ls := res.([]string)
		return ls[0], ls[1], nil
	}

	fk, ok, err := c._splitKey(k)
	if err != nil {
		return "", "", err
	}

	ls := []string{fk, ok}
	c.cache.Store(ck, ls)

	return fk, ok, nil
}

// getfilepath: get file key
func (c cfg) getfilepath(k string) (string, error) {
	fk, _, err := c.splitKey(k)
	if err != nil {
		return "", fmt.Errorf("error getting file path: %v", err)
	}
	return fk, nil
}

// getik: get inner cfg key
func (c cfg) getik(k string) string {
	_, ok, err := c.splitKey(k)
	if err != nil {
		panic("error getting ik")
	}
	return ok

}

// isnk: checks if the key is nested
func (c cfg) isnk(k string) bool {
	ik := c.getik(k)
	return len(ik) > 0
}

func (c cfg) loadFile(k string) (*objs.Obj, error) {

	die := func(f string, a ...any) (*objs.Obj, error) {
		return nil, fmt.Errorf(f, a...)
	}

	file, err := c.getfilepath(k)
	if err != nil {
		return die("error getting file key: %v", err)
	}

	obj, err := objs.FromFile(file)
	if err != nil {
		return die("error loading file: %v", err)
	}

	return obj, nil
}

func CFGGetStrict[T any](c cfg, arg ...string) T {
	k := strings.Join(arg, ".")
	typ := typeString[T]() // ← Get string name for type T

	key := getck(typ, k)
	if res, ok := c.cache.Load(key); ok {
		return res.(T)
	}

	obj, err := c.loadFile(k)
	if err != nil {
		panic(err)
	}

	var v any
	switch any(*new(T)).(type) {
	case int, string, bool, float64, map[string]string, map[string]any, []any, []string:
		v = objs.ObjGetStrict[T](obj, c.getik(k))
	case *objs.Obj:
		// TODO, send this validation down
		if strs.IsEmpty(c.getik(k)) {
			v = obj
		} else {
			v = obj.GetObj(c.getik(k))
		}
	default:
		panic("unsupported type")
	}

	c.cache.Store(key, v)
	return v.(T)
}

func (c cfg) GetBool(arg ...string) bool {
	return CFGGetStrict[bool](c, arg...)
}

func (c cfg) GetFloat64(arg ...string) float64 {
	return CFGGetStrict[float64](c, arg...)
}

func (c cfg) GetInt(arg ...string) int {
	return CFGGetStrict[int](c, arg...)
}

func (c cfg) GetString(arg ...string) string {
	return CFGGetStrict[string](c, arg...)
}

func (c cfg) GetStringMap(arg ...string) map[string]string {
	return CFGGetStrict[map[string]string](c, arg...)
}

func (c cfg) GetMap(arg ...string) map[string]any {
	return CFGGetStrict[map[string]any](c, arg...)
}

func (c cfg) GetSlice(arg ...string) []any {
	return CFGGetStrict[[]any](c, arg...)
}

func (c cfg) GetStringSlice(arg ...string) []string {
	return CFGGetStrict[[]string](c, arg...)
}

func (c cfg) GetObj(arg ...string) *objs.Obj {
	return CFGGetStrict[*objs.Obj](c, arg...)
}

func (c cfg) GetAs(as any, arg ...string) {

	k := strings.Join(arg, ".")

	cfg, err := c.loadFile(k)
	if err != nil {
		panic(err)
	}

	cfg.GetAs(c.getik(k), as)
}

func (c cfg) get(k string) (any, error) {

	var (
		cfg *objs.Obj
		res any
	)

	cfg, err := c.loadFile(k)
	if err != nil {
		return nil, err
	}

	if c.isnk(k) {
		res = cfg.Get(c.getik(k))
	} else {
		res = cfg.Data
	}

	return res, nil
}

func (c cfg) Get(arg ...string) any {
	k := strings.Join(arg, ".")
	ck := getck("any", k)
	if res, ok := c.cache.Load(ck); ok {
		return res
	}

	v, err := c.get(k)
	if err != nil {
		panic(err)
	}
	c.cache.Store(ck, v)
	return v
}
