package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/mem"
)

type P struct {
	Command
}

// P[<m>] -- display contents of memory slot
// m: memory slot (default: 0)
func NewP(memory mem.Memory) P {
	return P{New(memory)}
}

func (c P) Run(parameters []string) int {
	fmt.Println(c.ReadAll(parameters))
	return 0
}

func (c P) ReadAll(parameters []string) string {
	return string(c.Memory.Get(c.ToInt("slot (m)", parameters, 0, 0)))
}
