package set

import (
	"encoding/json"
	"sync"
)

type Set[T comparable] struct {
	mu   sync.RWMutex
	data map[T]struct{}
}

// New returns a new instance of Set.
func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Set adds a value to the set.
func (s *Set[T]) Set(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[value] = struct{}{}
}

// Has checks if a value exists in the set.
func (s *Set[T]) Has(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.data[value]
	return ok
}

// Delete removes a value from the set.
func (s *Set[T]) Delete(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, value)
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// Ls returns all values in the set as a slice.
func (s *Set[T]) Ls() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	values := make([]T, 0, len(s.data))
	for k := range s.data {
		values = append(values, k)
	}
	return values
}

// Union returns a new set containing all elements from both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()

	s.mu.RLock()
	for k := range s.data {
		result.Set(k)
	}
	s.mu.RUnlock()

	other.mu.RLock()
	for k := range other.data {
		result.Set(k)
	}
	other.mu.RUnlock()

	return result
}

// Intersection returns a new set containing only the elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()

	s.mu.RLock()
	other.mu.RLock()
	for k := range s.data {
		if _, ok := other.data[k]; ok {
			result.Set(k)
		}
	}
	other.mu.RUnlock()
	s.mu.RUnlock()

	return result
}

// Difference returns a new set containing elements in s but not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()

	s.mu.RLock()
	other.mu.RLock()
	for k := range s.data {
		if _, ok := other.data[k]; !ok {
			result.Set(k)
		}
	}
	other.mu.RUnlock()
	s.mu.RUnlock()

	return result
}

// IsSubset returns true if s is a subset of other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	s.mu.RLock()
	other.mu.RLock()
	defer s.mu.RUnlock()
	defer other.mu.RUnlock()

	for k := range s.data {
		if _, ok := other.data[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s is a superset of other.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// Equal returns true if both sets contain the same elements.
func (s *Set[T]) Equal(other *Set[T]) bool {
	s.mu.RLock()
	other.mu.RLock()
	defer s.mu.RUnlock()
	defer other.mu.RUnlock()

	if len(s.data) != len(other.data) {
		return false
	}
	for k := range s.data {
		if _, ok := other.data[k]; !ok {
			return false
		}
	}
	return true
}

// MarshalJSON encodes the set as a JSON array.
func (s *Set[T]) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return json.Marshal(s.Ls())
}

// UnmarshalJSON decodes a JSON array into the set.
func (s *Set[T]) UnmarshalJSON(b []byte) error {
	var items []T
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[T]struct{})
	for _, item := range items {
		s.data[item] = struct{}{}
	}

	return nil
}
