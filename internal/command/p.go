package command

import (
	"fmt"
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
	return P{New(memory)}
}

func (self P) Run(parameters []string) int {
	fmt.Fprintln(os.Stderr, self.Get(parameters))
	return 0
}

func (self P) Get(parameters []string) string {
	m := self.Range("m (memory slot)", parameters, 0, 0, 0, mem.MaxSlots-1)
	return string(self.Memory.Slots[m])
}
