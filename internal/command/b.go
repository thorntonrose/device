package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/mem"
)

type B struct {
	Command
}

func NewB(memory *mem.Memory) Runner {
	return B{New(memory)}
}

func (b B) Run(parameters Parameters) {
	b.Set(&b.Memory.Source, b.Buffer(parameters[0]))
	b.Set(&b.Memory.Destination, b.Buffer(parameters[1]))
}

func (b B) Set(curr *int, buf int) {
	if buf > 0 && buf < 9 {
		*curr = buf
	}

	b.Reset(*curr, buf)
}

func (b B) Reset(curr int, buf int) {
	if buf == 9 {
		b.Memory.Reset(curr)
	}
}

func (b B) Buffer(parameter int) int {
	if (parameter >= 0 && parameter <= mem.MaxBuffers) || parameter == 9 {
		return parameter
	}

	panic(fmt.Sprintf("invalid parameter: %d", parameter))
}
