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
	b := self.Range("b (buffer)", parameters, 0, 0, 0, mem.MaxBuffers)
	log.Printf("V.Run: b: %d\n", b)

	fmt.Fprintln(os.Stderr, string(*self.Memory.Buffers[Value(b, self.Memory.Destination)]))
	return
}
