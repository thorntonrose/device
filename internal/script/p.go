package script

import (
	"fmt"
	"log"
	"os"

	"github.com/thorntonrose/device/internal/mem"
)

// P[m] -- display contents of memory slot
//
// m: memory slot (default: 0)
type P struct {
	Command
}

func NewP(memory *mem.Memory) P {
	return P{NewCommand(memory)}
}

func (self P) Run(parameters []string) int {
	log.Printf("P.Run: %v\n", parameters)
	m := self.Range("m (memory slot)", parameters, 0, 0, 0, mem.MaxSlots-1)
	fmt.Fprintf(os.Stderr, "%03d:%s\n", m, string(self.Memory.Slots[m]))

	return 0
}
