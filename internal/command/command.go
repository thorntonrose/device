package command

import "github.com/thorntonrose/device/internal/mem"

type Command struct {
	Memory *mem.Memory
	Runner
}

func New(memory *mem.Memory) Command {
	return Command{Memory: memory}
}

type Runner interface {
	Run(parameters Parameters)
}

//-----------------------------------------------------------------------------

type Parameters [mem.MaxVariables]int
