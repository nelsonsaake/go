package lfs

// lookup: returns the first file found in the given list of paths
func Lookup(paths ...string) string {
	return NewBuilder().SetPaths(paths...).Lookup()
}
