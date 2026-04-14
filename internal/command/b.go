package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/mem"
)

// B<b1>.<b2> -- Set source and destination buffers
// b1: source buffer (0 - <max-buffers> | 9, default: 0)
// b2: destination buffer (0 - <max-buffers> | 9, default: 0)
//
// 0 = no change
// 1 - <max-buffers> = set buffer; reset pointer
// 9 = reset pointer
type B struct {
	Command
}

func NewB(memory *mem.Memory) B {
	return B{New(memory)}
}

func (c B) Run(parameters []string) int {
	b1 := c.Buffer("b1 (source)", parameters, 0)
	b2 := c.Buffer("b2 (destination)", parameters, 1)
	c.Set(&c.Memory.Source, b1)
	c.Set(&c.Memory.Destination, b2)

	return 0
}

func (c B) Set(buf *int, b int) {
	if b > 0 && b < 9 {
		*buf = b
	}

	c.Reset(*buf, b)
}

func (c B) Reset(buf int, b int) {
	if b == 9 {
		c.Memory.Reset(buf)
	}
}

func (c B) Buffer(name string, parameters []string, index int) int {
	buf := c.Int(name, parameters, index, 0)

	if (buf >= 0 && buf <= mem.MaxBuffers) || buf == 9 {
		return buf
	}

	panic(fmt.Sprintf("invalid: %s: %d", name, buf))
}
