package command

import (
	"log"

	"github.com/thorntonrose/device/device/mem"
)

// *M -- return from subroutine (script)
type StarM struct {
	Command
}

func NewStarM(memory *mem.Memory) StarM {
	return StarM{New(memory)}
}

func (StarM) Run(parameters []string) (skip int) {
	log.Println("*M")
	return mem.MaxBufferSize + 1
}
