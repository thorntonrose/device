package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

// G -- clear destination buffer
type G struct {
	Command
}

func NewG(memory *mem.Memory) G {
	return G{New(memory)}
}

func (g G) Run(_ []string) int {
	g.Memory.Clear(g.Memory.Destination)
	return 0
}
