package bufs

import (
	"log"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/mem"
)

// G -- clear destination buffer
type G struct {
	command.Command
}

func NewG(memory *mem.Memory) G {
	return G{command.New(memory)}
}

func (self G) Run(_ []string) int {
	log.Printf("G.Run: clear: %d\n", self.Memory.Destination)
	self.Memory.Clear(self.Memory.Destination)

	return 0
}
