package script

import (
	"github.com/thorntonrose/device/internal/mem"
)

// G -- clear destination buffer
type G struct {
	Command
}

func NewG(memory *mem.Memory) G {
	return G{NewCommand(memory)}
}

func (g G) Run(_ []string) int {
	g.Memory.Clear(g.Memory.Destination)
	return 0
}
