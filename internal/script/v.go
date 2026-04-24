package script

import (
	"fmt"
	"os"

	"github.com/thorntonrose/device/internal/mem"
)

// V<b> -- display contents of buffer
// b: buffer (0 - <max-buffers>, default: 0)
//
// 0 = destination buffer
// 1 - <max-buffers> = buffer number
type V struct {
	Command
}

func NewV(memory *mem.Memory) V {
	return V{NewCommand(memory)}
}

func (c V) Run(parameters []string) int {
	fmt.Fprintln(os.Stderr, c.ReadAll(parameters))
	return 0
}

func (c V) ReadAll(parameters []string) string {
	b := c.Int("b (buffer)", parameters, 0, mem.Transmit)
	// ???: Should start from read pointer, unless buffer = destination buffer.
	return string(c.Memory.ReadAll(b, 0, 0))
}
