package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// B[b1.b2] -- Set source and destination buffers
//
// b1: source buffer (0 - 5 | 9, default: 0); 0 = no change, 1 - 5 = set buffer and reset pointer; 9 = reset pointer
// b2: destination buffer (0 - 5 | 9, default: 0); 0 = no change, 1 - 5 = set buffer
type B struct {
	Command
}

func NewB(memory *mem.Memory) B {
	return B{New(memory)}
}

func (self B) Run(parameters []string) int {
	b1 := self.Code("b1 (source)", parameters, 0, 0, []int{0, 1, 2, 3, 4, 5, 9})
	b2 := self.Code("b2 (destination)", parameters, 1, 0, []int{0, 1, 2, 3, 4, 5})
	self.SetSource(b1)
	self.SetDestination(b2)

	return 0
}

func (self B) SetSource(b int) {
	if b > 0 && b != 9 {
		self.Memory.Source = b
	}

	self.Memory.Reset(self.Memory.Source)
}

func (self B) SetDestination(b int) {
	if b > 0 {
		self.Memory.Destination = b
	}
}
