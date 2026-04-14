package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// A<m>[.<n>] -- append data to destination buffer
// m: memory slot
// n: number of commands to skip if slot is empty
type A struct {
	Command
}

func NewA(memory mem.Memory) A {
	return A{Command{Memory: memory}}
}

func (c A) Run(parameters []string) (skip int) {
	if data := c.Memory.Get(c.ToInt("slot (m)", parameters, 0, 0)); len(data) > 0 {
		c.Memory.WriteAll(c.Memory.Destination, data)
		return 0
	}

	return c.ToInt("skip (n)", parameters, 1, 0)
}
