package script

import (
	"log"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

type N struct {
	Command
}

// N[n] -- write newlines to destination buffer
//
// n: number of newlines (default: 1)
func NewN(memory *mem.Memory) N {
	return N{NewCommand(memory)}
}

func (self N) Run(parameters []string) int {
	log.Printf("N.Run: %v\n", parameters)
	n := self.Positive("n (newlines)", parameters, 0, 1)
	etc.Times(n, func(_ int) { self.Memory.Write(self.Memory.Destination, byte('\n')) })

	return 0
}
