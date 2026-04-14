package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// O[<m>] -- move read pointer
// n: number of characters to move (positive or negative, default: 1)
type O struct {
	Command
}

func NewO(memory mem.Memory) O {
	return O{New(memory)}
}

func (c O) Run(parameters []string) int {
	c.Memory.Get(mem.Pointers)[c.Memory.Source] += byte(c.ToInt("number to move (n)", parameters, 0, 1))
	return 0
}
