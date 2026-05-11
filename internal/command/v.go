package command

import (
	"fmt"
	"log"
	"os"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// V[b] -- display contents of buffer (without moving pointer)
type V struct {
	Command
}

func NewV(memory *mem.Memory) V {
	return V{New(memory)}
}

func (self V) Run(parameters []string) (skip int) {
	log.Printf("V.Run: %v\n", parameters)
	b := self.Range("b (buffer)", parameters, 0, 0, 0, mem.MaxBuffers)
	fmt.Fprintln(os.Stderr, string(*self.Memory.Buffers[Value(b, self.Memory.Destination)]))

	return
}
