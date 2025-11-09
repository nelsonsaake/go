package dsc

import (
	"sync"

	"github.com/nelsonsaake/go/strs"
)

// Subject holds information about a code generation target.
type Subject struct {
	id   string
	data map[string]string
	rsrc []string
}

// cache for memoized New() calls
var genDesCache sync.Map

// New creates or retrieves a memoized GenDes instance.
func New() *Subject {

	id := strs.UUID()

	if val, ok := genDesCache.Load(id); ok {
		return val.(*Subject)
	}

	g := &Subject{
		id:   id,
		data: map[string]string{},
	}

	genDesCache.Store(id, g)
	return g
}

func (d *Subject) WithValue(k, v string) *Subject {
	d.data[k] = v
	return d
}

func (d *Subject) WithData(v map[string]string) *Subject {
	d.data = v
	return d
}

func (d *Subject) WithID(v ...string) *Subject {
	d.id = ID(v...)
	return d
}

func (d *Subject) WithRSrc(rsrc ...string) *Subject {
	d.rsrc = rsrc
	return d
}

// Get returns the value of a field by key (case-insensitive).
func (s Subject) Get(key string) string {

	res, ok := s.data[key]
	if ok {
		return res
	}

	res, ok = s.data[strs.ToLower(key)]
	if ok {
		return res
	}

	return "<nil>"
}

// ID returns a unique identifier for this GenDes (used in Is lookups).
func (s Subject) ID() string {
	return s.id
}

// Is checks if the GenDes instance has a certain boolean property tag.
// Delegates to a global Is(...) function (assumed to be implemented elsewhere).
func (s Subject) Is(predicate string) bool {
	rsrc := append(s.rsrc, predicate)
	return Is(s, s.ID(), rsrc...)
}

func (s Subject) Tr(target string) any {
	return Tr(s, s.ID(), target, s.rsrc...)
}
