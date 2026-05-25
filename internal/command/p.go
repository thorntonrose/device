package command

import (
	"fmt"
	"log"
	"os"

	"github.com/thorntonrose/device/internal/mem"
)

// P[m] -- display contents of memory slot
type P struct {
	Command
}

func NewP(memory *mem.Memory) P {
	return P{New(memory)}
}

func (self P) Run(parameters []string) (skip int) {
	m := self.Range("m (memory slot)", parameters, 0, 0, 0, mem.MaxSlots-1)
	log.Printf("P: m: %d\n", m)

	fmt.Fprintf(os.Stderr, "%03d:%s\n", m, string(self.Memory.Slots[m]))
	return
}
