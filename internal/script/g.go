package script

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// G -- clear destination buffer
type G struct {
	Command
}

func NewG(memory *mem.Memory) G {
	return G{NewCommand(memory)}
}

func (self G) Run(_ []string) int {
	log.Printf("G.Run: clear: %d\n", self.Memory.Destination)
	self.Memory.Clear(self.Memory.Destination)

	return 0
}
