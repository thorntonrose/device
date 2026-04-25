package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// A[m.s] -- append data to destination buffer
//
// m: memory slot (default: 0)
// s: commands to skip if slot is empty (default: 0)
type A struct {
	Command
}

func NewA(memory *mem.Memory) A {
	return A{New(memory)}
}

func (self A) Run(parameters []string) (skip int) {
	m := self.Range("m (memory slot)", parameters, 0, 0, 0, mem.MaxSlots-1)
	s := self.Int("s (skip)", parameters, 1, 0)
	self.WriteAll(m)

	return s
}

func (self A) WriteAll(m int) {
	if data := self.Memory.Slots[m]; len(data) > 0 {
		self.Memory.WriteAll(self.Memory.Destination, data)
	}
}
