package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/mem/buf"
)

// X[n.c] -- copy source buffer to destination buffer (moving pointer)
type X struct {
	Command
}

func NewX(memory *mem.Memory) X {
	return X{New(memory)}
}

func (self X) Run(parameters []string) (skip int) {
	n := self.NonNegative("n (characters)", parameters, 0, 0)
	c := self.Range("c (stop character)", parameters, 1, 0, 0, 255)
	log.Printf("X: n: %d, c: %d\n", n, c)

	self.Copy(n, byte(c))
	return
}

func (self X) Copy(n int, c byte) {
	data := buf.ReadAll(self.Memory, self.Memory.Source, n, byte(c))
	log.Printf("Copy: dest: %d, data: %v\n", self.Memory.Destination, data)
	buf.WriteAll(self.Memory, self.Memory.Destination, data)
}
