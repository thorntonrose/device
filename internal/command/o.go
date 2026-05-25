package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// O[n] -- move source buffer pointer
type O struct {
	Command
}

func NewO(memory *mem.Memory) O {
	return O{New(memory)}
}

func (self O) Run(parameters []string) (skip int) {
	n := self.Int("n (number to move)", parameters, 0, 1)
	log.Printf("O: n: %d\n", n)

	self.Memory.Pointers[self.Memory.Source] += n
	return
}
