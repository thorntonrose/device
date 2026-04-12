package command

import "github.com/thorntonrose/device/internal/mem"

type X struct {
	Command
}

func NewX(memory *mem.Memory) Runner {
	return X{New(memory)}
}

func (x X) Run(parameters Parameters) {
	x.Memory.WriteAll(x.Memory.Destination, x.Memory.ReadAll(x.Memory.Source, parameters[0], byte(parameters[1])))
}
