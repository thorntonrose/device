package script

import (
	"fmt"
	"os"

	"github.com/thorntonrose/device/internal/mem"
)

type P struct {
	Command
}

// P<m> -- display contents of memory slot
// m: memory slot (default: 0)
func NewP(memory *mem.Memory) P {
	return P{NewCommand(memory)}
}

func (c P) Run(parameters []string) int {
	fmt.Fprintln(os.Stderr, c.Get(parameters))
	return 0
}

func (c P) Get(parameters []string) string {
	m := c.Int("m (memory slot)", parameters, 0, 0)
	return string(c.Memory.Get(m))
}
