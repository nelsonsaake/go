package paths

// Path: resolve from the root of the app, if we find the root
// otherwise, we return it as it
func (p *pfs) Path(ls ...string) string {
	return p.Join(p.root, ls)
}
