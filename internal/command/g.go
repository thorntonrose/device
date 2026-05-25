package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/mem/buf"
)

// G -- clear destination buffer
type G struct {
	Command
}

func NewG(memory *mem.Memory) G {
	return G{New(memory)}
}

func (self G) Run(_ []string) (skip int) {
	log.Printf("G: dest: %d\n", self.Memory.Destination)
	buf.Clear(self.Memory, self.Memory.Destination)

	return
}
