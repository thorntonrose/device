package script

import "github.com/thorntonrose/device/internal/mem"

type StarO struct {
	Command
}

func NewStarO(memory *mem.Memory) StarO {
	return StarO{NewCommand(memory)}
}
