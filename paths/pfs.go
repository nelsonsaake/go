package paths

type FS struct {
	root string
	unix bool
}

func (p *FS) Root() string {
	return p.root
}

func New(fsRoot string) *FS {
	return &FS{root: fsRoot}
}

func NewUnix(fsRoot string) *FS {
	return &FS{root: fsRoot, unix: true}
}
