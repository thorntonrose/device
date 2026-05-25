package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/mem/buf"
)

// R[c] -- append constant value to destination buffer
type R struct {
	Command
}

func NewR(memory *mem.Memory) R {
	return R{New(memory)}
}

func (self R) Run(parameters []string) (skip int) {
	c := self.Bytes("c (constant value)", parameters, 0, []byte{'\t'})
	log.Printf("R: c: %v (%s)\n", c, string(c))

	buf.WriteAll(self.Memory, self.Memory.Destination, c)
	return
}
