package command

import (
	"log"
	"strings"

	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
)

// H[s.c] -- search for string in source buffer
type H struct {
	Command
}

func NewH(memory *mem.Memory) H {
	return H{New(memory)}
}

func (self H) Run(parameters []string) (skip int) {
	s := self.Int("s (skip)", parameters, 0, 0)
	c := self.Bytes("c (string)", parameters, 1, []byte{'\t'})
	log.Printf("H: s: %d, c: %v (%s)\n", s, c, string(c))

	return If(self.Find(c) == -1, s, 0)
}

func (self H) Find(c []byte) (index int) {
	buf := *self.Memory.Buffers[self.Memory.Source]
	ptr := &self.Memory.Pointers[self.Memory.Source]
	index = strings.Index(string(buf[*ptr:]), string(c))
	*ptr += If(index == -1, 0, index)

	return
}
