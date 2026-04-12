package command

import (
	"github.com/thorntonrose/device/internal/mem"
)

type G struct {
	Command
}

func NewG(memory *mem.Memory) Runner {
	return G{New(memory)}
}

func (g G) Run(_ Parameters) {
	g.Memory.Clear(g.Memory.Destination)
}
