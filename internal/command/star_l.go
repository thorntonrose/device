package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

type Script interface {
	Run(slotNum int) int
}

// *L[m] -- call subroutine (script)
type StarL struct {
	Command
	Script Script
}

func NewStarL(memory *mem.Memory, script Script) StarL {
	return StarL{New(memory), script}
}

func (self StarL) Run(parameters []string) (skip int) {
	log.Printf("StarL.Run: %v\n", parameters)
	m := self.Int("m (memory slot)", parameters, 0, 0)
	self.Script.Run(m)

	return
}
