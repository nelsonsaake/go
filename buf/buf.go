package buf

import (
	"bytes"
	"fmt"
)

type Buf struct {
	w bytes.Buffer
}

func (b *Buf) Ln(a ...any) {
	b.w.WriteString(fmt.Sprintln(a...))
}

func New() (b *Buf) {
	return &Buf{
		w: bytes.Buffer{},
	}
}

func (b *Buf) String() string {
	return b.w.String()
}
