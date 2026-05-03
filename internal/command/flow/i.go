package flow

import (
	"log"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// I[s.a.c] -- compare constant value <c> to source buffer and skip
//
// s: commands to skip (positive or negative) if <a> is true
// a: comparison code (0 - 4, default: 0); 0 = true, 1 = equal, 2 = not equal, 3 = less than, 4 = greater than
// c: constant value to compare (0 - 255 | '<text>', default: ASCII FS [28])
//
// Note:
// 1. Skipping past the beginning or end of the script slot terminates the script.
// 2. If the source buffer starting from the pointer is empty, skip always occurs.
type I struct {
	command.Command
}

func NewI(memory *mem.Memory) I {
	return I{command.New(memory)}
}

func (self I) Run(parameters []string) (skip int) {
	log.Printf("I.Run: %v\n", parameters)
	s := self.Int("s (skip)", parameters, 0, 0)
	a := self.Code("a (comparison)", parameters, 1, 0, []int{0, 1, 2, 3, 4})
	c := string(self.Bytes("c (constant)", parameters, 2, []byte{28}))

	return etc.If(self.Compare(self.Data(c), a, c), s, 0)
}

func (self I) Compare(s string, a int, c string) bool {
	result := (a == 0) || (a == 1 && s == c) || (a == 2 && s != c) || (a == 3 && s < c) || (a == 4 && s > c)
	log.Printf("I.Compare: %t\n", result)

	return result
}

func (self I) Data(c string) string {
	buf := string(*self.Memory.Buffers[self.Memory.Source])
	ptr := self.Memory.Pointers[self.Memory.Source]

	data := etc.If(ptr+len(c) <= len(buf), func() string { return buf[ptr : ptr+len(c)] },
		func() string { return "" })()

	log.Printf("I.Data: %s\n", data)
	return data
}
