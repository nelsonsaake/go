package register

import (
	"fmt"
	"sync"

	"github.com/nelsonsaake/go/strs"
)

// Generic Register type
type Register[T any] struct {
	index map[string]T
	mu    sync.RWMutex
}

// Global storage for singletons
var (
	defaultRegistries sync.Map               // type-based
	namedRegistries   = make(map[string]any) // name-based
	namedMu           sync.Mutex
)

// New returns a singleton per type
func New[T any]() *Register[T] {
	// unique key per type
	key := typeKey[T]()
	if reg, ok := defaultRegistries.Load(key); ok {
		return reg.(*Register[T])
	}
	r := &Register[T]{index: make(map[string]T)}
	defaultRegistries.Store(key, r)
	return r
}

// Named returns a singleton per type+name
func Named[T any](name string) *Register[T] {
	key := typeKey[T]() + ":" + name

	namedMu.Lock()
	defer namedMu.Unlock()

	if reg, ok := namedRegistries[key]; ok {
		return reg.(*Register[T])
	}

	r := &Register[T]{index: make(map[string]T)}
	namedRegistries[key] = r
	return r
}

// Register a new item
func (r *Register[T]) Register(name string, value T) {
	if strs.IsNotSnakeCase(name) {
		panic("registry name must be snake cased: " + name + " given")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.index[name] = value
}

// Get an item by name
func (r *Register[T]) Get(name string) (T, bool) {
	if strs.IsNotSnakeCase(name) {
		panic("registry name must be snake cased: " + name + " given")
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok := r.index[name]
	return value, ok
}

// GetAll returns a copy of all registered items as a map
func (r *Register[T]) GetAll() map[string]T {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make(map[string]T, len(r.index))
	for k, v := range r.index {
		result[k] = v
	}
	return result
}

// ---- helpers ----

// typeKey generates a unique string key per type
func typeKey[T any]() string {
	var zero T
	return fmt.Sprintf("%T", zero)
}
