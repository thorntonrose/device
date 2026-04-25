package command

import (
	"fmt"
	"os"

	"github.com/thorntonrose/device/internal/mem"
)

// V[b] -- display contents of buffer (without moving pointer)
//
// b: buffer (0 - <max-buffers>, default: 0); 0 = destination buffer
type V struct {
	Command
}

func NewV(memory *mem.Memory) V {
	return V{New(memory)}
}

func (self V) Run(parameters []string) int {
	fmt.Fprintln(os.Stderr, self.Get(parameters))
	return 0
}

func (self V) Get(parameters []string) string {
	b := self.Range("b (buffer)", parameters, 0, self.Memory.Destination, 0, mem.MaxBuffers)
	return string(*self.Memory.Buffers[b])
}
