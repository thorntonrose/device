package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// G -- clear destination buffer
type G struct {
	Command
}

func NewG(memory *mem.Memory) G {
	return G{New(memory)}
}

func (self G) Run(_ []string) (skip int) {
	log.Printf("G.Run: clear: %d\n", self.Memory.Destination)
	self.Memory.Clear(self.Memory.Destination)

	return
}
