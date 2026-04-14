package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/mem"
)

// V[<b>] -- display contents of buffer
// b: buffer (0 - <max-buffers>, default: 0)
//
// 0 = destination buffer
// 1 - <max-buffers> = buffer number
type V struct {
	Command
}

func NewV(memory mem.Memory) V {
	return V{New(memory)}
}

func (c V) Run(parameters []string) int {
	fmt.Println(c.ReadAll(parameters))
	return 0
}

func (c V) ReadAll(parameters []string) string {
	buf := c.ToInt("buffer (b)", parameters, 0, 0)
	return string(c.Memory.ReadAll(buf, 0, 0))
}
