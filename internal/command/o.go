package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// O[n] -- move source buffer pointer
//
// n: number of characters to move (positive or negative, default: 1)
type O struct {
	Command
}

func NewO(memory *mem.Memory) O {
	return O{New(memory)}
}

func (self O) Run(parameters []string) int {
	n := self.Int("n (number to move)", parameters, 0, 1)
	self.Memory.Pointers[self.Memory.Source] += n

	return 0
}
