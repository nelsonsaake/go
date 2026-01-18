package paths

type pfs struct {
	root string
	unix bool
}

func (p *pfs) Root() string {
	return p.root
}

func New(fsRoot string) *pfs {
	return &pfs{root: fsRoot}
}

func NewUnix(fsRoot string) *pfs {
	return &pfs{root: fsRoot, unix: true}
}
