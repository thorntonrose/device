package bufs

import (
	"log"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/mem"
)

// O[n] -- move source buffer pointer
//
// n: characters to move (positive or negative, default: 1)
type O struct {
	command.Command
}

func NewO(memory *mem.Memory) O {
	return O{command.New(memory)}
}

func (self O) Run(parameters []string) int {
	log.Printf("O.Run: %v\n", parameters)
	n := self.Int("n (number to move)", parameters, 0, 1)

	self.Memory.Pointers[self.Memory.Source] += n
	return 0
}
