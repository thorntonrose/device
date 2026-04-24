package script

import (
	"github.com/thorntonrose/device/internal/mem"
)

// O<n> -- move source buffer pointer
// n: number of characters to move (positive or negative, default: 1)
type O struct {
	Command
}

func NewO(memory *mem.Memory) O {
	return O{NewCommand(memory)}
}

func (c O) Run(parameters []string) int {
	n := c.Int("n (number to move)", parameters, 0, 1)
	c.Memory.Pointers[c.Memory.Source] += n

	return 0
}
