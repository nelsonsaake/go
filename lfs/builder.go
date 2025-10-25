package lfs

import (
	"io/fs"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/ufs"
)

type LookupBuilder struct {
	fs    fs.FS // Use standard library fs.FS interface
	paths []string
}

// NewBuilder returns a new LookupBuilder instance
func NewBuilder() *LookupBuilder {
	return &LookupBuilder{}
}

func (b *LookupBuilder) WithFS(v fs.FS) *LookupBuilder {
	b.fs = v
	return b
}

// Add appends single path
func (b *LookupBuilder) Add(path ...string) *LookupBuilder {
	b.paths = append(b.paths, afs.Path(path...))
	return b
}

// Add multiple paths
func (b *LookupBuilder) SetPaths(path ...string) *LookupBuilder {
	b.paths = append(b.paths, path...)
	return b
}

// Lookup performs the actual lookup and returns the first match
func (b *LookupBuilder) Lookup() string {
	for _, v := range b.paths {
		exists, _ := ufs.Exists(v)
		if exists {
			return v
		}
	}
	return ""
}
