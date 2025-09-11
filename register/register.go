package register

import (
	"fmt"
	"sync"

	"github.com/nelsonsaake/go/strs"
)

// Generic Registry type
type Registry[T any] struct {
	index map[string]T
	mu    sync.RWMutex
}

// Global storage for singletons
var (
	defaultRegistries sync.Map               // type-based
	namedRegistries   = make(map[string]any) // name-based
	namedMu           sync.Mutex
)

// NewRegistry returns a singleton per type
func NewRegistry[T any]() *Registry[T] {
	// unique key per type
	key := typeKey[T]()
	if reg, ok := defaultRegistries.Load(key); ok {
		return reg.(*Registry[T])
	}
	r := &Registry[T]{index: make(map[string]T)}
	defaultRegistries.Store(key, r)
	return r
}

// NewNamedRegistry returns a singleton per type+name
func NewNamedRegistry[T any](name string) *Registry[T] {
	key := typeKey[T]() + ":" + name

	namedMu.Lock()
	defer namedMu.Unlock()

	if reg, ok := namedRegistries[key]; ok {
		return reg.(*Registry[T])
	}

	r := &Registry[T]{index: make(map[string]T)}
	namedRegistries[key] = r
	return r
}

// Register a new item
func (r *Registry[T]) Register(name string, value T) {
	if strs.IsNotSnakeCase(name) {
		panic("registry name must be snake cased: " + name + " given")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.index[name] = value
}

// Get an item by name
func (r *Registry[T]) Get(name string) (T, bool) {
	if strs.IsNotSnakeCase(name) {
		panic("registry name must be snake cased: " + name + " given")
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok := r.index[name]
	return value, ok
}

// ---- helpers ----

// typeKey generates a unique string key per type
func typeKey[T any]() string {
	var zero T
	return fmt.Sprintf("%T", zero)
}
