package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// *M -- return from subroutine (script)
type StarM struct {
	Command
}

func NewStarM(memory *mem.Memory) StarM {
	return StarM{New(memory)}
}

func (c StarM) Run(parameters []string) (skip int) {
	log.Println("*M.Run")
	return mem.MaxBufferSize + 1
}
