package jfs

import "path/filepath"

// File: json file
type File struct {
	path string
}

func (j *File) Write(data any) error {
	return Marshal(j.path, data)
}

func (j *File) Read(v any) error {
	return Unmarshal(j.path, v)
}

func New(path ...string) *File {
	return &File{filepath.Join(path...)}
}
