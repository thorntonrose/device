package command

import (
	"fmt"
	"io"
	"log"
	"os"

	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
)

var Display io.Writer = os.Stderr

// V[b] -- display contents of buffer (without moving pointer)
type V struct {
	Command
}

func NewV(memory *mem.Memory) V {
	return V{New(memory)}
}

func (self V) Run(parameters []string) (skip int) {
	b := self.Range("b (buffer)", parameters, 0, 0, 0, mem.MaxBuffers)
	log.Printf("V: b: %d\n", b)

	fmt.Fprintln(Display, string(*self.Memory.Buffers[Value(b, self.Memory.Destination)]))
	return
}
