package command

import "github.com/thorntonrose/device/internal/mem"

type StarM struct {
	Command
}

func NewStarM(memory *mem.Memory) StarM {
	return StarM{New(memory)}
}

func (c StarM) Run(parameters []string) (skip int) {
	return mem.MaxBufferSize + 1
}
