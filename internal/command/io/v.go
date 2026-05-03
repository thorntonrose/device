package io

import (
	"fmt"
	"log"
	"os"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// V[b] -- display contents of buffer (without moving pointer)
//
// b: buffer (0 - <max-buffers>, default: 0) 0 = destination buffer
type V struct {
	command.Command
}

func NewV(memory *mem.Memory) V {
	return V{command.New(memory)}
}

func (self V) Run(parameters []string) int {
	log.Printf("V.Run: %v\n", parameters)
	b := self.Range("b (buffer)", parameters, 0, 0, 0, mem.MaxBuffers)

	fmt.Fprintln(os.Stderr, string(*self.Memory.Buffers[etc.Value(b, self.Memory.Destination)]))
	return 0
}
