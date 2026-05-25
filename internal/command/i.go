package command

import (
	"log"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// I[s.a.c] -- compare constant value to source buffer and skip
type I struct {
	Command
}

func NewI(memory *mem.Memory) I {
	return I{New(memory)}
}

func (self I) Run(parameters []string) (skip int) {
	s := self.Int("s (skip)", parameters, 0, 0)
	a := self.Code("a (comparison)", parameters, 1, 0, []int{0, 1, 2, 3, 4})
	c := string(self.Bytes("c (constant)", parameters, 2, []byte{'\t'}))
	log.Printf("I: s: %d, a: %d, c: %s\n", s, a, c)

	return If(self.Compare(self.Data(c), a, c), s, 0)
}

func (self I) Compare(s string, a int, c string) (result bool) {
	result = (a == 0) || (a == 1 && s == c) || (a == 2 && s != c) || (a == 3 && s < c) || (a == 4 && s > c)
	log.Printf("Compare: %t\n", result)

	return
}

func (self I) Data(c string) (data string) {
	buf := string(*self.Memory.Buffers[self.Memory.Source])
	ptr := self.Memory.Pointers[self.Memory.Source]
	data = If(ptr+len(c) <= len(buf), func() string { return buf[ptr : ptr+len(c)] }, func() string { return "" })()
	log.Printf("Data: %v (%s)\n", []byte(data), data)

	return
}
