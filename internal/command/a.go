package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// A<m>.<s> -- append data to destination buffer
// m: memory slot (default: 0)
// s: commands to skip if slot is empty (default: 0)
type A struct {
	Command
}

func NewA(memory *mem.Memory) A {
	return A{New(memory)}
}

func (c A) Run(parameters []string) (skip int) {
	m := c.Int("m (memory slot)", parameters, 0, 0)
	s := c.Int("s (skip)", parameters, 1, 0)
	c.WriteAll(m)

	return s
}

func (c A) WriteAll(m int) {
	if data := c.Memory.Get(m); len(data) > 0 {
		c.Memory.WriteAll(c.Memory.Destination, data)
	}
}
