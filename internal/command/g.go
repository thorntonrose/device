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

func (self G) Run(_ []string) int {
	self.Memory.Clear(self.Memory.Destination)
	return 0
}
